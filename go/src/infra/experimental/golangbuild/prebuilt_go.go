// Copyright 2023 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/client/cmd/swarming/swarmingimpl"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/luciexe/build"
)

// prebuiltGo represents a mapping between a Go toolchain version and the prebuilt
// GOROOT for that toolchain in CAS.
type prebuiltGo struct {
	// ID represents the toolchain build target. Specifically, it takes the form: $GOOS-$GOARCH-$commit.
	ID string `gae:"$id"`

	// CASDigest is the digest of the prebuilt toolchain in CAS.
	//
	// Note: this is optimistic and might be stale. CAS may throw away
	// the prebuilt toolchain at any time, but usually keeps it around for
	// at least a couple days.
	CASDigest string

	// Extra and unrecognized fields will be loaded without issues, but not saved.
	_ datastore.PropertyMap `gae:"-,extra"`
}

func (p *prebuiltGo) String() string {
	return fmt.Sprintf("%q -> %q", p.ID, p.CASDigest)
}

func casInstanceFromEnv() (string, error) {
	// Obtain the instance from SWARMING_SERVER like recipes do.
	//
	// It may be a bit weird to import this variable from a command
	// implementation, but other CLI executables in LUCI do it too.
	// Also it means if this changes, it's likely it'll get noticed
	// by whoever changes it.
	server := os.Getenv(swarmingimpl.ServerEnvVar)
	if server == "" {
		return "", fmt.Errorf("no CAS instance found")
	}
	u, err := url.Parse(server)
	if err != nil {
		return "", fmt.Errorf("%q is not a URL: %v", swarmingimpl.ServerEnvVar, err)
	}
	inst, found := strings.CutSuffix(u.Host, ".appspot.com")
	if !found {
		return "", fmt.Errorf("%q is not an appspot.com URL", swarmingimpl.ServerEnvVar)
	}
	return inst, nil
}

func checkForPrebuiltGo(ctx context.Context, src *sourceSpec) (digest string, err error) {
	step, ctx := build.StartStep(ctx, "check for prebuilt go")
	defer func() {
		// Any failure in this function is an infrastructure failure.
		err = build.AttachStatus(err, bbpb.Status_INFRA_FAILURE, nil)
		step.End(err)
	}()

	id, err := src.prebuiltID()
	if err != nil {
		return "", err
	}
	tc := &prebuiltGo{
		ID: id,
	}
	switch err := datastore.Get(ctx, tc); {
	case err == datastore.ErrNoSuchEntity:
		return "", nil
	case err != nil:
		return "", err
	}
	_, err = io.WriteString(step.Log("mapping"), tc.String())
	if err != nil {
		return "", err
	}
	return tc.CASDigest, nil
}

func fetchGoFromCAS(ctx context.Context, spec *buildSpec, digest, goroot string) (ok bool, err error) {
	step, ctx := build.StartStep(ctx, "fetch prebuilt go")
	defer func() {
		// Any failure in this function is an infrastructure failure.
		err = build.AttachStatus(err, bbpb.Status_INFRA_FAILURE, nil)
		step.End(err)
	}()

	// Create a file to write out structured results.
	//
	// We're passing this to a command via filename but don't
	// close it yet; we'll be able to read from it after that
	// command exits.
	jsonDump, err := os.CreateTemp("", "golangbuild-cas-json")
	if err != nil {
		return false, err
	}
	defer jsonDump.Close()

	// Run 'cas download'.
	cmd := spec.toolCmd(ctx,
		"cas", "download",
		"-cas-instance", spec.casInstance,
		"-dir", goroot,
		"-digest", digest,
		"-dump-json", jsonDump.Name(),
	)
	if err := runCommandAsStep(ctx, "cas download", cmd, true); err != nil {
		var dlr struct {
			Result string `json:"result"`
		}
		if err := json.NewDecoder(jsonDump).Decode(&dlr); err != nil {
			return false, err
		}
		if dlr.Result == "digest_invalid" {
			// The prebuilt toolchain isn't available in CAS anymore.
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func uploadGoToCAS(ctx context.Context, spec *buildSpec, src *sourceSpec, goroot string) (err error) {
	step, ctx := build.StartStep(ctx, "upload prebuilt go")
	defer func() {
		// Any failure in this function is an infrastructure failure.
		err = build.AttachStatus(err, bbpb.Status_INFRA_FAILURE, nil)
		step.End(err)
	}()

	// Collect the paths that we'll be archiving.
	gorootEntries, err := os.ReadDir(goroot)
	if err != nil {
		return err
	}
	var pathArgs []string
	for _, entry := range gorootEntries {
		// No reason to keep the .git directory.
		if entry.Name() == ".git" {
			continue
		}
		pathArgs = append(pathArgs, "-paths", fmt.Sprintf("%s:%s", goroot, entry.Name()))
	}

	// Create a file to write out the digest.
	//
	// We're passing this to a command via filename but don't
	// close it yet; we'll be able to read from it after that
	// command exits.
	digestFile, err := os.CreateTemp("", "golangbuild-cas-digest")
	if err != nil {
		return err
	}
	defer digestFile.Close()

	// Run 'cas archive'.
	args := []string{
		"archive",
		"-cas-instance", spec.casInstance,
		"-dump-digest", digestFile.Name(),
	}
	cmd := spec.toolCmd(ctx, "cas", append(args, pathArgs...)...)
	if err := runCommandAsStep(ctx, "cas archive", cmd, true); err != nil {
		return err
	}

	// Read the digest output.
	output, err := io.ReadAll(digestFile)
	if err != nil {
		return err
	}

	// Get the prebuilt ID.
	id, err := src.prebuiltID()
	if err != nil {
		return err
	}

	// Update the datastore with the digest.
	tc := &prebuiltGo{
		ID:        id,
		CASDigest: strings.TrimSpace(string(output)),
	}
	_, err = io.WriteString(step.Log("mapping"), tc.String())
	if err != nil {
		return err
	}
	return datastore.Put(ctx, tc)
}