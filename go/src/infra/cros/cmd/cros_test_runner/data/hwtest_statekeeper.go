// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package data

import (
	vmlabapi "infra/libs/vmlab/api"

	testapi "go.chromium.org/chromiumos/config/go/test/api"
	artifactpb "go.chromium.org/chromiumos/config/go/test/artifact"
	labapi "go.chromium.org/chromiumos/config/go/test/lab/api"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/skylab_test_runner"
	"go.chromium.org/luci/luciexe/build"
	"google.golang.org/protobuf/types/known/anypb"

	"infra/cros/cmd/common_lib/interfaces"
	"infra/cros/cmd/common_lib/tools/crostoolrunner"
	"infra/cros/dutstate"
)

// HwTestStateKeeper represents all the data hw test execution flow requires.
type HwTestStateKeeper struct {
	interfaces.StateKeeper

	// Build related
	BuildState *build.State

	// Set from input
	CftTestRequest *skylab_test_runner.CFTTestRequest

	// Dut related
	HostName         string
	DutTopology      *labapi.DutTopology
	DutServerAddress *labapi.IpEndpoint
	CurrentDutState  dutstate.State
	// Only when DUT is a VM
	DutVmGceImage   *vmlabapi.GceImage
	DutVm           *vmlabapi.VmInstance
	LeaseVMResponse *testapi.LeaseVMResponse

	// Provision related
	InstallMetadata *anypb.Any
	ProvisionResp   *testapi.InstallResponse

	// Test related
	TestArgs      *testapi.AutotestExecutionMetadata
	TestResponses *testapi.CrosTestResponse

	// Publish related
	GcsUrl              string
	StainlessUrl        string
	TesthausUrl         string
	GcsPublishSrcDir    string
	CurrentInvocationId string
	TkoPublishSrcDir    string
	CpconPublishSrcDir  string
	TestResultForRdb    *artifactpb.TestResult

	// Build related
	SkylabResult *skylab_test_runner.Result

	// Tools and their related dependencies
	Ctr                   *crostoolrunner.CrosToolRunner
	DockerKeyFileLocation string
}