// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package validation

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/chromiumos/config/go/test/api"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestLeaseVMValidate(t *testing.T) {
	Convey("LeaseVMRequest Validate", t, func() {
		Convey("Valid request - successful path", func() {
			req := &api.LeaseVMRequest{
				HostReqs: &api.VMRequirements{
					GceImage:       "test-image",
					GceRegion:      "test-region",
					GceProject:     "test-project",
					GceMachineType: "test-machine-type",
					GceDiskSize:    100,
				},
			}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldBeNil)
		})
		Convey("Invalid request - missing host requirements", func() {
			req := &api.LeaseVMRequest{}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "VM requirements must be set.")
		})
		Convey("Invalid request - missing image", func() {
			req := &api.LeaseVMRequest{
				HostReqs: &api.VMRequirements{
					GceRegion:      "test-region",
					GceProject:     "test-project",
					GceMachineType: "test-machine-type",
					GceDiskSize:    100,
				},
			}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE image must be set.")
		})
		Convey("Invalid request - missing region", func() {
			req := &api.LeaseVMRequest{
				HostReqs: &api.VMRequirements{
					GceImage:       "test-image",
					GceProject:     "test-project",
					GceMachineType: "test-machine-type",
					GceDiskSize:    100,
				},
			}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE region (zone) must be set.")
		})
		Convey("Invalid request - missing project", func() {
			req := &api.LeaseVMRequest{
				HostReqs: &api.VMRequirements{
					GceImage:       "test-image",
					GceRegion:      "test-region",
					GceMachineType: "test-machine-type",
					GceDiskSize:    100,
				},
			}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE project must be set.")
		})
		Convey("Invalid request - missing machine type", func() {
			req := &api.LeaseVMRequest{
				HostReqs: &api.VMRequirements{
					GceImage:    "test-image",
					GceRegion:   "test-region",
					GceProject:  "test-project",
					GceDiskSize: 100,
				},
			}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE machine type must be set.")
		})
		Convey("Invalid request - missing disk size", func() {
			req := &api.LeaseVMRequest{
				HostReqs: &api.VMRequirements{
					GceImage:       "test-image",
					GceRegion:      "test-region",
					GceProject:     "test-project",
					GceMachineType: "test-machine-type",
				},
			}
			err := ValidateLeaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE machine disk size must be set (in GB).")
		})
	})
}

func TestReleaseVMValidate(t *testing.T) {
	Convey("ReleaseVMRequest Validate", t, func() {
		Convey("Valid request - successful path", func() {
			req := &api.ReleaseVMRequest{
				LeaseId:    "test-lease-id",
				GceProject: "test-project",
				GceRegion:  "test-region",
			}
			err := ValidateReleaseVMRequest(req)
			So(err, ShouldBeNil)
		})
		Convey("Invalid request - missing lease id", func() {
			req := &api.ReleaseVMRequest{
				GceProject: "test-project",
				GceRegion:  "test-region",
			}
			err := ValidateReleaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "Lease ID must be set.")
		})
		Convey("Invalid request - missing project", func() {
			req := &api.ReleaseVMRequest{
				LeaseId:   "test-lease-id",
				GceRegion: "test-region",
			}
			err := ValidateReleaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE project must be set.")
		})
		Convey("Invalid request - missing region", func() {
			req := &api.ReleaseVMRequest{
				LeaseId:    "test-lease-id",
				GceProject: "test-project",
			}
			err := ValidateReleaseVMRequest(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE region (zone) must be set.")
		})
	})
}

func TestVMRequirementsValidate(t *testing.T) {
	Convey("VMRequirements Validate", t, func() {
		Convey("Valid request - successful path", func() {
			req := &api.VMRequirements{
				GceImage:       "test-image",
				GceRegion:      "test-region",
				GceProject:     "test-project",
				GceMachineType: "test-machine-type",
				GceDiskSize:    100,
			}
			err := ValidateVMRequirements(req)
			So(err, ShouldBeNil)
		})
		Convey("Invalid request - missing image", func() {
			req := &api.VMRequirements{
				GceRegion:      "test-region",
				GceProject:     "test-project",
				GceMachineType: "test-machine-type",
				GceDiskSize:    100,
			}
			err := ValidateVMRequirements(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE image must be set.")
		})
		Convey("Invalid request - missing region", func() {
			req := &api.VMRequirements{
				GceImage:       "test-image",
				GceProject:     "test-project",
				GceMachineType: "test-machine-type",
				GceDiskSize:    100,
			}
			err := ValidateVMRequirements(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE region (zone) must be set.")
		})
		Convey("Invalid request - missing project", func() {
			req := &api.VMRequirements{
				GceImage:       "test-image",
				GceRegion:      "test-region",
				GceMachineType: "test-machine-type",
				GceDiskSize:    100,
			}
			err := ValidateVMRequirements(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE project must be set.")
		})
		Convey("Invalid request - missing machine type", func() {
			req := &api.VMRequirements{
				GceImage:    "test-image",
				GceRegion:   "test-region",
				GceProject:  "test-project",
				GceDiskSize: 100,
			}
			err := ValidateVMRequirements(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE machine type must be set.")
		})
		Convey("Invalid request - missing disk size", func() {
			req := &api.VMRequirements{
				GceImage:       "test-image",
				GceRegion:      "test-region",
				GceProject:     "test-project",
				GceMachineType: "test-machine-type",
			}
			err := ValidateVMRequirements(req)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "GCE machine disk size must be set (in GB).")
		})
	})
}

func TestValidateLeaseParent(t *testing.T) {
	Convey("Validate LeaseParent", t, func() {
		Convey("Valid regex - successful path; only project", func() {
			err := ValidateLeaseParent("projects/test-project")
			So(err, ShouldBeNil)
		})
		Convey("Valid regex - successful path; project and zone", func() {
			err := ValidateLeaseParent("projects/test-project/zones/test-zone")
			So(err, ShouldBeNil)
		})
		Convey("Valid regex - error; no project", func() {
			err := ValidateLeaseParent("projects/")
			So(err, ShouldErrLike, "parent must be in the format `projects/${project}` or `projects/${project}/zones/${zone}`")
		})
		Convey("Valid regex - error; extra string", func() {
			err := ValidateLeaseParent("projects/test-project/123")
			So(err, ShouldErrLike, "parent must be in the format `projects/${project}` or `projects/${project}/zones/${zone}`")
		})
		Convey("Valid regex - error; typo in zone", func() {
			err := ValidateLeaseParent("projects/test-project/zone/fail-zone")
			So(err, ShouldErrLike, "parent must be in the format `projects/${project}` or `projects/${project}/zones/${zone}`")
		})
	})
}
