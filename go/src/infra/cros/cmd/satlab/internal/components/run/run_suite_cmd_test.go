// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package run

import (
	"context"
	"testing"

	"github.com/googleapis/gax-go/v2"
	buildbucketpb "go.chromium.org/luci/buildbucket/proto"
	moblabpb "google.golang.org/genproto/googleapis/chromeos/moblab/v1beta1"

	"infra/cros/cmd/satlab/internal/pkg/google.golang.org/google/chromeos/moblab"
)

// FakeMoblabClient is a mock Moblab API client that returns hardcoded data
type FakeMoblabClient struct{}

func (f *FakeMoblabClient) StageBuild(ctx context.Context, req *moblabpb.StageBuildRequest, opts ...gax.CallOption) (*moblab.StageBuildOperation, error) {

	return &moblab.StageBuildOperation{}, nil
}

func (f *FakeMoblabClient) CheckBuildStageStatus(context.Context, *moblabpb.CheckBuildStageStatusRequest, ...gax.CallOption) (*moblabpb.CheckBuildStageStatusResponse, error) {

	name := "buildTargets/octopus/models/bobba/builds/1234.0.0/artifacts/chromeos-moblab-peng-staging"
	return &moblabpb.CheckBuildStageStatusResponse{
		IsBuildStaged:       true,
		StagedBuildArtifact: &moblabpb.BuildArtifact{Name: name},
	}, nil
}

// FakeBuildbucketClient is a mock Buildbucket client that returns hardcoded data
type FakeBuildbucketClient struct {
	badData bool
}

func (f *FakeBuildbucketClient) ScheduleCTPBuild(context.Context) (*buildbucketpb.Build, error) {
	if f.badData {
		return &buildbucketpb.Build{}, nil
	}

	return &buildbucketpb.Build{
		Id: 0000,
		Builder: &buildbucketpb.BuilderID{
			Project: "project",
			Bucket:  "bucket",
			Builder: "builder",
		},
	}, nil
}

// TestRunSuite tests the innerRun function of our command with fake Moblab and Buildbucket clients
// It tests input entirely, partially, and not at all user given
func TestRunSuite(t *testing.T) {
	t.Parallel()

	type test struct {
		inputCommand *runSuite
	}

	tests := []test{
		{
			&runSuite{
				runSuiteFlags: runSuiteFlags{
					suite:     "rlz",
					board:     "zork",
					model:     "gumboz",
					milestone: "111",
					build:     "15329.6.0"},
			},
		},
	}

	for _, tc := range tests {
		fakeMoblabClient := FakeMoblabClient{}
		fakeBuildbucketClient := FakeBuildbucketClient{badData: false}
		bucket := "chromeos-distributed-fleet-s4p"
		err := tc.inputCommand.innerRunWithClients(context.Background(), &fakeMoblabClient, &fakeBuildbucketClient, bucket)
		if err != nil {
			t.Errorf("Unexpected err: %v", err)
		}
	}
}