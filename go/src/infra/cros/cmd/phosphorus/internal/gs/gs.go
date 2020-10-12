// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package gs exports helpers to upload log data to Google Storage.
package gs

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"go.chromium.org/luci/common/sync/parallel"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/gcloud/gs"
	gcgs "go.chromium.org/luci/common/gcloud/gs"
	"go.chromium.org/luci/common/logging"
)

// DirWriter exposes methods to write a local directory to Google Storage.
type DirWriter struct {
	client               gsClient
	maxConcurrentUploads int
}

// gsClient is a Google Storage client.
//
// This interface is a subset of the gcgs.Client interface.
type gsClient interface {
	NewWriter(p gcgs.Path) (gs.Writer, error)
}

// NewDirWriter creates an object which can write a directory and its subdirectories to the given Google Storage path
func NewDirWriter(client gsClient) *DirWriter {
	return &DirWriter{
		client:               client,
		maxConcurrentUploads: maxConcurrentUploads,
	}
}

func verifyPaths(localPath string, gsPath string) error {
	problems := []string{}
	if _, err := os.Stat(localPath); err != nil {
		problems = append(problems, fmt.Sprintf("invalid local path (%s)", localPath))
	}
	if _, err := url.Parse(gsPath); err != nil {
		problems = append(problems, fmt.Sprintf("invalid GS path (%s)", gsPath))
	}
	if len(problems) > 0 {
		return errors.Reason("path errors: %s", strings.Join(problems, ", ")).Err()
	}
	return nil
}

const maxConcurrentUploads = 10

type file struct {
	Src  string
	Dest gcgs.Path
	Info os.FileInfo
}

// WriteDir writes a local directory to Google Storage.
//
// If ctx is canceled, WriteDir() returns after completing in-flight uploads,
// skipping remaining contents of the directory and returns ctx.Err().
func (w *DirWriter) WriteDir(ctx context.Context, srcDir string, dstDir gcgs.Path) error {
	logging.Debugf(ctx, "Writing %s and subtree to %s.", srcDir, dstDir)
	if err := verifyPaths(srcDir, string(dstDir)); err != nil {
		return err
	}

	var merr errors.MultiError
	files := []*file{}
	if err := filepath.Walk(srcDir, func(src string, info os.FileInfo, err error) error {
		// Continue walking the directory tree on errors so that we upload as
		// many files as possible.
		if err != nil {
			merr = append(merr, errors.Annotate(err, "list files to upload").Err())
			return nil
		}
		relPath, err := filepath.Rel(srcDir, src)
		if err != nil {
			merr = append(merr, errors.Annotate(err, "writing from %s to %s", src, dstDir).Err())
			return nil
		}
		files = append(files, &file{
			Src:  src,
			Dest: dstDir.Concat(relPath),
			Info: info,
		})
		return nil
	}); err != nil {
		panic(fmt.Sprintf("Directory walk leaked error: %s", err))
	}

	var terr error
	err := parallel.WorkPool(w.maxConcurrentUploads, func(items chan<- func() error) {
		for _, f := range files {
			// Create a loop-local variable for capture in the lambda.
			f := f
			item := func() error {
				return w.writeOne(ctx, f.Src, f.Dest, f.Info)
			}
			select {
			case items <- item:
				continue
			case <-ctx.Done():
				// terr is captured.
				terr = ctx.Err()
				break
			}
		}
	})
	if err != nil {
		merr = append(merr, err)
	}
	if terr != nil {
		merr = append(merr, err)
	}
	if len(merr) > 0 {
		return errors.Annotate(merr, "writing dir %s to %s", srcDir, dstDir).Err()
	}
	return nil
}

func (w *DirWriter) writeOne(ctx context.Context, src string, dest gcgs.Path, info os.FileInfo) error {
	if info.IsDir() {
		return nil
	}
	if skip, reason := shouldSkipUpload(info); skip {
		logging.Debugf(ctx, "Skipped %s because: %s.", src, reason)
		return nil
	}

	f, err := os.Open(src)
	if err != nil {
		return errors.Annotate(err, "writing from %s to %s", src, dest).Err()
	}
	defer f.Close()

	writer, err := w.client.NewWriter(dest)
	if err != nil {
		return errors.Annotate(err, "writing from %s to %s", src, dest).Err()
	}
	// Ignore errors as we may have already closed writer by the time this runs.
	defer writer.Close()

	bs := make([]byte, info.Size())
	if _, err = f.Read(bs); err != nil {
		return errors.Annotate(err, "writing from %s to %s", src, dest).Err()
	}
	n, err := writer.Write(bs)
	if err != nil {
		return errors.Annotate(err, "writing from %s to %s", src, dest).Err()
	}
	if int64(n) != info.Size() {
		return errors.Reason("length written to %s does not match source file size", dest).Err()
	}
	err = writer.Close()
	if err != nil {
		return errors.Annotate(err, "writer for %s failed to close", dest).Err()
	}
	return nil
}

// shouldSkipUpload determines if a particular file should be skipped.
//
// Also returns a reason for skipping the file.
func shouldSkipUpload(i os.FileInfo) (bool, string) {
	if i.Mode()&os.ModeType == 0 {
		return false, ""
	}

	switch {
	case i.Mode()&os.ModeSymlink == os.ModeSymlink:
		return true, "file is a symlink"
	case i.Mode()&os.ModeDevice == os.ModeDevice:
		return true, "file is a device"
	case i.Mode()&os.ModeNamedPipe == os.ModeNamedPipe:
		return true, "file is a named pipe"
	case i.Mode()&os.ModeSocket == os.ModeSocket:
		return true, "file is a unix domain socket"
	case i.Mode()&os.ModeIrregular == os.ModeIrregular:
		return true, "file is an irregular file of unknown type"
	default:
		return true, "file is a non-file of unknown type"
	}
}
