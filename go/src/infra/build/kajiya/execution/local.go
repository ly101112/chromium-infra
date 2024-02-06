// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package execution

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/bazelbuild/remote-apis-sdks/go/pkg/digest"
	repb "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"infra/build/kajiya/blobstore"
)

// Executor is a local executor that executes actions on the local machine.
// It uses a ContentAddressableStorage to fetch all required inputs for the action
// into a sandbox directory, executes the action in that sandbox directory, and
// then uploads the output files and directories to the CAS after the action has finished.
type Executor struct {
	cas         *blobstore.ContentAddressableStorage
	images      *ImageRepository
	nsjailPath  string
	sandboxBase string
}

// New creates a new Executor.
// baseDir is a directory used to store temporary files required during execution, such as
// sandbox directories.
// cas is the ContentAddressableStorage to use for fetching and uploading blobs.
func New(baseDir string, cas *blobstore.ContentAddressableStorage) (*Executor, error) {
	if baseDir == "" {
		return nil, fmt.Errorf("baseDir must be set")
	}

	if cas == nil {
		return nil, fmt.Errorf("cas must be set")
	}

	// Create the data directory if it doesn't exist.
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %q: %w", baseDir, err)
	}

	images, err := NewImageRepository(filepath.Join(baseDir, "images"))
	if err != nil {
		return nil, fmt.Errorf("failed to create image repository: %w", err)
	}

	sandboxBase := filepath.Join(baseDir, "tmp")
	if err := os.MkdirAll(sandboxBase, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %q: %w", sandboxBase, err)
	}

	// Try to find nsjail in the PATH.
	nsjailPath := ""
	if runtime.GOOS == "linux" {
		nsjailPath, err = exec.LookPath("nsjail")
		if err != nil {
			return nil, fmt.Errorf("🚨 required tool 'nsjail' not found in PATH")
		}
	}

	return &Executor{
		cas:         cas,
		images:      images,
		nsjailPath:  nsjailPath,
		sandboxBase: sandboxBase,
	}, nil
}

// Execute executes the given action and returns the result.
func (e *Executor) Execute(action *repb.Action) (*repb.ActionResult, error) {
	var missingBlobs []digest.Digest

	// Get the command from the CAS.
	cmd, err := e.getCommand(action.CommandDigest)
	if err != nil {
		if os.IsNotExist(err) {
			missingBlobs = append(missingBlobs, digest.NewFromProtoUnvalidated(action.CommandDigest))
		} else {
			return nil, err
		}
	}

	// Get the input root from the CAS.
	inputRoot, err := e.getDirectory(action.InputRootDigest)
	if err != nil {
		if os.IsNotExist(err) {
			missingBlobs = append(missingBlobs, digest.NewFromProtoUnvalidated(action.InputRootDigest))
			return nil, e.formatMissingBlobsError(missingBlobs)
		} else {
			return nil, err
		}
	}

	// Build a sandbox directory for the action.
	sandboxDir, err := os.MkdirTemp(e.sandboxBase, "*")
	if err != nil {
		return nil, fmt.Errorf("failed to create sandbox directory: %w", err)
	}
	defer e.deleteSandbox(sandboxDir)

	// Materialize the input root in the sandbox directory.
	mb, err := e.materializeDirectory(sandboxDir, inputRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to materialize input root: %w", err)
	}
	missingBlobs = append(missingBlobs, mb...)

	// If there were any missing blobs, we fail early and return the list to the client.
	if len(missingBlobs) > 0 {
		return nil, e.formatMissingBlobsError(missingBlobs)
	}

	// If a working directory was specified, verify that it exists.
	workDir := sandboxDir
	if cmd.WorkingDirectory != "" {
		if !filepath.IsLocal(cmd.WorkingDirectory) {
			return nil, fmt.Errorf("working directory %q points outside of input root", cmd.WorkingDirectory)
		}
		workDir = filepath.Join(sandboxDir, cmd.WorkingDirectory)
		if err := os.MkdirAll(workDir, 0755); err != nil {
			return nil, fmt.Errorf("could not create working directory: %v", err)
		}
	}

	// Create the directories required by all output files and directories.
	outputPaths, err := e.createOutputPaths(cmd, workDir)
	if err != nil {
		return nil, err
	}

	// Execute the command.
	actionResult, err := e.executeCommand(sandboxDir, action, cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %v", err)
	}

	// Save stdout and stderr to the CAS and update their digests in the action result.
	if err := e.saveStdOutErr(actionResult); err != nil {
		return nil, err
	}

	// Go through all output files and directories and upload them to the CAS.
	for _, outputPath := range outputPaths {
		joinedPath := filepath.Join(workDir, outputPath)
		fi, err := os.Stat(joinedPath)
		if err != nil {
			if os.IsNotExist(err) {
				// Ignore non-existing output files.
				continue
			}
			return nil, fmt.Errorf("failed to stat output path %q: %v", outputPath, err)
		}
		if fi.IsDir() {
			// Upload the directory to the CAS.
			dirs, err := e.buildMerkleTree(joinedPath)
			if err != nil {
				return nil, fmt.Errorf("failed to build merkle tree for %q: %v", outputPath, err)
			}

			tree := repb.Tree{
				Root: dirs[0],
			}
			if len(dirs) > 1 {
				tree.Children = dirs[1:]
			}
			treeBytes, err := proto.Marshal(&tree)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal tree: %v", err)
			}
			d, err := e.cas.Put(treeBytes)
			if err != nil {
				return nil, fmt.Errorf("failed to upload tree to CAS: %v", err)
			}

			actionResult.OutputDirectories = append(actionResult.OutputDirectories, &repb.OutputDirectory{
				Path:                  outputPath,
				TreeDigest:            d.ToProto(),
				IsTopologicallySorted: false,
			})
		} else {
			// Upload the file to the CAS.
			d, err := digest.NewFromFile(joinedPath)
			if err != nil {
				return nil, fmt.Errorf("failed to compute digest of file %q: %v", outputPath, err)
			}
			if err := e.cas.Adopt(d, joinedPath); err != nil {
				return nil, fmt.Errorf("failed to upload file %q to CAS: %v", outputPath, err)
			}

			actionResult.OutputFiles = append(actionResult.OutputFiles, &repb.OutputFile{
				Path:         outputPath,
				Digest:       d.ToProto(),
				IsExecutable: fi.Mode()&0111 != 0,
			})
		}
	}

	return actionResult, nil
}

// createOutputPaths creates the directories required by all output files and directories.
// It transforms and returns the list of output paths so that they're relative to our current working directory.
func (e *Executor) createOutputPaths(cmd *repb.Command, workDir string) (outputPaths []string, err error) {
	if cmd.OutputPaths != nil {
		// REAPI v2.1+
		outputPaths = cmd.OutputPaths
	} else {
		// REAPI v2.0
		outputPaths = make([]string, 0, len(cmd.OutputFiles)+len(cmd.OutputDirectories))
		outputPaths = append(outputPaths, cmd.OutputFiles...)
		outputPaths = append(outputPaths, cmd.OutputDirectories...)
	}
	for _, outputPath := range outputPaths {
		// We need to create the parent directories of the output path, because the command
		// may not create them itself.
		if err := os.MkdirAll(filepath.Join(workDir, filepath.Dir(outputPath)), 0755); err != nil {
			return nil, fmt.Errorf("failed to create parent directories for %q: %v", outputPath, err)
		}
	}
	return outputPaths, nil
}

// saveStdOutErr saves stdout and stderr to the CAS and returns the updated action result.
func (e *Executor) saveStdOutErr(actionResult *repb.ActionResult) error {
	d, err := e.cas.Put(actionResult.StdoutRaw)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to put stdout into CAS: %v", err)
	}
	actionResult.StdoutDigest = d.ToProto()

	d, err = e.cas.Put(actionResult.StderrRaw)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to put stderr into CAS: %v", err)
	}
	actionResult.StderrDigest = d.ToProto()

	// Servers are not required to inline stdout and stderr, so we just set them to nil.
	// The client can just fetch them from the CAS if it needs them.
	actionResult.StdoutRaw = nil
	actionResult.StderrRaw = nil

	return nil
}

func (e *Executor) getDirectory(d *repb.Digest) (*repb.Directory, error) {
	dirDigest, err := digest.NewFromProto(d)
	if err != nil {
		return nil, fmt.Errorf("failed to parse directory digest: %v", err)
	}
	dirBytes, err := e.cas.Get(dirDigest)
	if err != nil {
		return nil, fmt.Errorf("failed to get directory from CAS: %v", err)
	}
	dir := &repb.Directory{}
	if err := proto.Unmarshal(dirBytes, dir); err != nil {
		return nil, fmt.Errorf("failed to unmarshal directory: %v", err)
	}
	return dir, nil
}

func (e *Executor) getCommand(d *repb.Digest) (*repb.Command, error) {
	cmdDigest, err := digest.NewFromProto(d)
	if err != nil {
		return nil, fmt.Errorf("failed to parse command digest: %v", err)
	}
	cmdBytes, err := e.cas.Get(cmdDigest)
	if err != nil {
		return nil, fmt.Errorf("failed to get command from CAS: %v", err)
	}
	cmd := &repb.Command{}
	if err := proto.Unmarshal(cmdBytes, cmd); err != nil {
		return nil, fmt.Errorf("failed to unmarshal command: %v", err)
	}
	return cmd, nil
}

// materializeDirectory recursively materializes the given directory in the
// local filesystem. The directory itself is created at the given path, and
// all files and subdirectories are created under that path.
func (e *Executor) materializeDirectory(path string, d *repb.Directory) (missingBlobs []digest.Digest, err error) {
	// First, materialize all the input files in the directory.
	for _, fileNode := range d.Files {
		filePath := filepath.Join(path, fileNode.Name)
		err = e.materializeFile(filePath, fileNode)
		if err != nil {
			if os.IsNotExist(err) {
				missingBlobs = append(missingBlobs, digest.NewFromProtoUnvalidated(fileNode.Digest))
				continue
			}
			return nil, fmt.Errorf("failed to materialize file: %v", err)
		}
	}

	// Next, materialize all the subdirectories.
	for _, sdNode := range d.Directories {
		sdPath := filepath.Join(path, sdNode.Name)
		err = os.Mkdir(sdPath, 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to create subdirectory: %v", err)
		}

		sd, err := e.getDirectory(sdNode.Digest)
		if err != nil {
			if os.IsNotExist(err) {
				missingBlobs = append(missingBlobs, digest.NewFromProtoUnvalidated(sdNode.Digest))
				continue
			}
			return nil, fmt.Errorf("failed to get subdirectory: %v", err)
		}

		sdMissingBlobs, err := e.materializeDirectory(sdPath, sd)
		missingBlobs = append(missingBlobs, sdMissingBlobs...)
		if err != nil {
			return nil, fmt.Errorf("failed to materialize subdirectory: %v", err)
		}
	}

	// Finally, set the directory properties. We have to do this after the files
	// have been materialized, because otherwise the mtime of the directory would
	// be updated to the current time.
	if d.NodeProperties != nil {
		if d.NodeProperties.Mtime != nil {
			time := d.NodeProperties.Mtime.AsTime()
			if err := os.Chtimes(path, time, time); err != nil {
				return nil, fmt.Errorf("failed to set mtime: %v", err)
			}
		}

		if d.NodeProperties.UnixMode != nil {
			if err := os.Chmod(path, os.FileMode(d.NodeProperties.UnixMode.Value)); err != nil {
				return nil, fmt.Errorf("failed to set mode: %v", err)
			}
		}
	}

	return missingBlobs, nil
}

// materializeFile downloads the given file from the CAS and writes it to the given path.
func (e *Executor) materializeFile(filePath string, fileNode *repb.FileNode) error {
	fileDigest, err := digest.NewFromProto(fileNode.Digest)
	if err != nil {
		return fmt.Errorf("failed to parse file digest: %v", err)
	}

	// Calculate the file permissions from all relevant fields.
	perm := os.FileMode(0644)
	if fileNode.NodeProperties != nil && fileNode.NodeProperties.UnixMode != nil {
		perm = os.FileMode(fileNode.NodeProperties.UnixMode.Value)
	}
	if fileNode.IsExecutable {
		perm |= 0111
	}

	if err := e.cas.LinkTo(fileDigest, filePath); err != nil {
		return fmt.Errorf("failed to link to file in CAS: %w", err)
	}

	if err := os.Chmod(filePath, perm); err != nil {
		return fmt.Errorf("failed to set mode: %w", err)
	}

	if fileNode.NodeProperties != nil && fileNode.NodeProperties.Mtime != nil {
		time := fileNode.NodeProperties.Mtime.AsTime()
		if err := os.Chtimes(filePath, time, time); err != nil {
			return fmt.Errorf("failed to set mtime: %v", err)
		}
	}

	return nil
}

// formatMissingBlobsError formats a list of missing blobs as a gRPC "FailedPrecondition" error
// as described in the Remote Execution API.
func (e *Executor) formatMissingBlobsError(blobs []digest.Digest) error {
	violations := make([]*errdetails.PreconditionFailure_Violation, 0, len(blobs))
	for _, b := range blobs {
		violations = append(violations, &errdetails.PreconditionFailure_Violation{
			Type:    "MISSING",
			Subject: fmt.Sprintf("blobs/%s/%d", b.Hash, b.Size),
		})
	}

	s, err := status.New(codes.FailedPrecondition, "missing blobs").WithDetails(&errdetails.PreconditionFailure{
		Violations: violations,
	})
	if err != nil {
		return fmt.Errorf("failed to create status: %v", err)
	}

	return s.Err()
}

// buildNsjailArgs builds the arguments for nsjail.
func (e *Executor) buildNsjailArgs(sandboxDir string, imageDir string, cmd *repb.Command) []string {
	// We use /mnt as the input root inside the sandbox. The exact path doesn't matter,
	// because all paths in REAPI are relative to the input root anyway, so the action
	// is relocatable and not tied to a specific input root.
	// TODO: Add support for overriding this via RBE's InputRootAbsolutePath feature?
	inputRoot := "/mnt"

	workDir := filepath.Join(inputRoot, cmd.WorkingDirectory)
	searchPath := []string{"/usr/local/sbin", "/usr/local/bin", "/usr/sbin", "/usr/bin", "/sbin", "/bin"}

	// If we're not using a container image, it means the action should be executed
	// on the host, so we need to set the image path to "/". Otherwise, nsjail will
	// run the command in a completely empty root filesystem, which will fail.
	if imageDir == "" {
		imageDir = "/"
	}

	args := []string{
		e.nsjailPath,
		"--quiet",
		"--chroot", imageDir,
		"--cwd", workDir,
		"--bindmount", fmt.Sprintf("%s:%s", sandboxDir, inputRoot),
		"--env", "HOME=" + workDir,
		"--env", "PATH=" + strings.Join(searchPath, ":"),
	}

	// We don't need to set the values of the environment variables here because we
	// already set them in the environment of the nsjail process.
	for _, env := range cmd.EnvironmentVariables {
		args = append(args, "--env", env.Name)
	}

	// Python and other tools use /dev/shm for shared memory, so we need to mount it.
	// Otherwise we get errors like this:
	// _multiprocessing.SemLock(kind, value, maxvalue)
	//   OSError: [Errno 38] Function not implemented
	args = append(args, "--tmpfsmount", "/dev/shm")

	// Some tools use /tmp for temporary files, so we need to mount it.
	// Otherwise we get errors like this:
	//   clang: error: unable to make temporary file: Read-only file system
	args = append(args, "--tmpfsmount", "/tmp")

	// nsjail applies rather strict resource limits by default, which can cause
	// some tools to fail. We disable them here for now, until we figure out
	// a set of limits that works for most tools.
	args = append(args, "--disable_rlimits")

	// nsjail doesn't look up argv[0] in PATH, so we need to resolve it ourselves
	// and pass it to nsjail via --exec_file.
	if !strings.ContainsAny(cmd.Arguments[0], "/") {
		for _, p := range searchPath {
			if _, err := os.Stat(filepath.Join(imageDir, p, cmd.Arguments[0])); err == nil {
				args = append(args, "--exec_file", filepath.Join(p, cmd.Arguments[0]))
				break
			}
		}
	}

	return append(args, "--")
}

// executeCommand runs cmd in the sandboxDir, which must already have been prepared by the caller.
// If we were able to execute the command, a valid ActionResult will be returned and error is nil.
// This includes the case where we ran the command, and it exited with an exit code != 0.
// However, if something went wrong during preparation or while spawning the process, an error is returned.
func (e *Executor) executeCommand(sandboxDir string, action *repb.Action, cmd *repb.Command) (*repb.ActionResult, error) {
	var args []string

	imageURL := e.images.ImageURL(action, cmd)

	if e.nsjailPath != "" {
		// Prepare the container image for the action.
		imageDir, err := e.images.FetchImage(imageURL)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch container image: %w", err)
		}

		// If nsjail is available, we use it to sandbox the command.
		args = e.buildNsjailArgs(sandboxDir, imageDir, cmd)
	} else {
		// Otherwise, we just run the command directly on the host.
		// However, we can only do this if the action doesn't require a container image.
		if imageURL != "" {
			return nil, fmt.Errorf("action requires container image, but nsjail is not available")
		}
	}
	args = append(args, cmd.Arguments...)

	c := exec.Command(args[0], args[1:]...)
	c.Dir = filepath.Join(sandboxDir, cmd.WorkingDirectory)

	for _, env := range cmd.EnvironmentVariables {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", env.Name, env.Value))
	}

	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	if err := c.Run(); err != nil {
		// ExitError just means that the command returned a non-zero exit code.
		// In that case we just set the ExitCode in the ActionResult to it.
		// However, other errors mean that something went wrong, and we need to
		// return them to the caller.
		if exitErr := (&exec.ExitError{}); !errors.As(err, &exitErr) {
			return nil, err
		}
	}

	return &repb.ActionResult{
		ExitCode:  int32(c.ProcessState.ExitCode()),
		StdoutRaw: stdout.Bytes(),
		StderrRaw: stderr.Bytes(),
	}, nil
}

// addDirectoryToTree recursively walks through the given directory and adds itself, all files and
// subdirectories to the given Tree.
func (e *Executor) buildMerkleTree(path string) (dirs []*repb.Directory, err error) {
	dir := &repb.Directory{}

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			subDirs, err := e.buildMerkleTree(filepath.Join(path, dirEntry.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to build merkle tree: %v", err)
			}
			d, err := digest.NewFromMessage(subDirs[0])
			if err != nil {
				return nil, fmt.Errorf("failed to get digest: %v", err)
			}
			dir.Directories = append(dir.Directories, &repb.DirectoryNode{
				Name:   dirEntry.Name(),
				Digest: d.ToProto(),
			})
			dirs = append(dirs, subDirs...)
		} else {
			d, err := digest.NewFromFile(filepath.Join(path, dirEntry.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to get digest: %v", err)
			}
			fi, err := dirEntry.Info()
			if err != nil {
				return nil, fmt.Errorf("failed to get file info: %v", err)
			}
			fileNode := &repb.FileNode{
				Name:         dirEntry.Name(),
				Digest:       d.ToProto(),
				IsExecutable: fi.Mode()&0111 != 0,
			}
			err = e.cas.Adopt(d, filepath.Join(path, dirEntry.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to move file into CAS: %v", err)
			}
			dir.Files = append(dir.Files, fileNode)
		}
	}

	dirBytes, err := proto.Marshal(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal directory: %v", err)
	}
	if _, err = e.cas.Put(dirBytes); err != nil {
		return nil, err
	}

	return append([]*repb.Directory{dir}, dirs...), nil
}

func (e *Executor) deleteSandbox(dir string) {
	// Walk through all directories inside the sandbox and give ourselves permission to delete
	// everything. This is necessary, because actions can create directories or files with wrong
	// permissions, causing os.RemoveAll to fail to delete them.
	_ = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			_ = os.Chmod(path, 0700)
		} else {
			// While we're here, attempt to remove the file. If it works, we save some
			// work later. If not, just ignore the error and continue, because
			// os.RemoveAll will take care of it.
			_ = os.Remove(path)
		}
		return nil
	})

	if err := os.RemoveAll(dir); err != nil {
		log.Printf("🚨 failed to remove sandbox: %v", err)
	}
}
