// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inventory

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/gae/service/datastore"
	"google.golang.org/protobuf/types/known/durationpb"

	fleet "infra/appengine/crosskylabadmin/api/fleet/v1"
	"infra/appengine/crosskylabadmin/internal/app/config"
	dsinventory "infra/appengine/crosskylabadmin/internal/app/frontend/datastore/inventory"
	dssv "infra/appengine/crosskylabadmin/internal/app/frontend/datastore/stableversion"
	"infra/libs/skylab/inventory"
)

var testLooksLikeFakeServoTests = []struct {
	in   string
	good bool
}{
	{``, false},
	{`dummy_host`, false},
	{`FAKE_SERVO_HOST`, false},
	{`chromeos6-row3-rack11-labstation`, true},
}

func TestLooksLikeFakeServo(t *testing.T) {
	for _, tt := range testLooksLikeFakeServoTests {
		name := fmt.Sprintf("(%s)", tt.in)
		t.Run(name, func(t *testing.T) {
			good := !looksLikeFakeServo(tt.in)
			if good != tt.good {
				t.Errorf("wanted: (%t) got: (%t)", tt.good, good)
			}
		})
	}
}

func TestGetStableVersion(t *testing.T) {
	Convey("Test GetStableVersion RPC -- stable versions exist", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()
		err := dssv.PutSingleCrosStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFaftStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-faft-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-firmware-version")
		So(err, ShouldBeNil)
		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				BuildTarget: "xxx-build-target",
				Model:       "xxx-model",
			},
		)
		So(err, ShouldBeNil)
		So(resp.CrosVersion, ShouldEqual, "xxx-cros-version")
		So(resp.FaftVersion, ShouldEqual, "xxx-faft-version")
		So(resp.FirmwareVersion, ShouldEqual, "xxx-firmware-version")
	})

	Convey("Test GetStableVersion RPC -- look up by hostname beaglebone", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()

		// use a fake beaglebone servo
		duts := []*inventory.DeviceUnderTest{
			{
				Common: &inventory.CommonDeviceSpecs{
					Attributes: []*inventory.KeyValue{
						{
							Key:   strptr("servo_host"),
							Value: strptr("xxx-beaglebone-servo"),
						},
					},
					Id:       strptr("xxx-id"),
					Hostname: strptr("xxx-hostname"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-model"),
						Board: strptr("xxx-build-target"),
					},
				},
			},
		}

		err := dsinventory.UpdateDUTs(ctx, duts)
		So(err, ShouldBeNil)

		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleCrosStableVersion(ctx, beagleboneServo, beagleboneServo, "xxx-beaglebone-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFaftStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-faft-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-firmware-version")
		So(err, ShouldBeNil)

		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				Hostname: "xxx-hostname",
			},
		)

		So(err, ShouldBeNil)
		So(resp.CrosVersion, ShouldEqual, "xxx-cros-version")
		So(resp.FaftVersion, ShouldEqual, "xxx-faft-version")
		So(resp.FirmwareVersion, ShouldEqual, "xxx-firmware-version")
		So(resp.ServoCrosVersion, ShouldEqual, "xxx-beaglebone-cros-version")
	})

	Convey("Test GetStableVersion RPC -- look up by hostname labstation", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()

		// use a fake labstation
		duts := []*inventory.DeviceUnderTest{
			{
				Common: &inventory.CommonDeviceSpecs{
					Attributes: []*inventory.KeyValue{
						{
							Key:   strptr("servo_host"),
							Value: strptr("xxx-labstation"),
						},
					},
					Id:       strptr("xxx-id"),
					Hostname: strptr("xxx-hostname"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-model"),
						Board: strptr("xxx-build-target"),
					},
				},
			},
			{
				Common: &inventory.CommonDeviceSpecs{
					Id:       strptr("xxx-labstation-id"),
					Hostname: strptr("xxx-labstation"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-labstation-model"),
						Board: strptr("xxx-labstation-board"),
					},
				},
			},
		}

		err := dsinventory.UpdateDUTs(ctx, duts)
		So(err, ShouldBeNil)

		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-labstation-board", "xxx-labstation-model", "xxx-labstation-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFaftStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-faft-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-firmware-version")
		So(err, ShouldBeNil)

		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				Hostname: "xxx-hostname",
			},
		)

		So(err, ShouldBeNil)
		So(resp.CrosVersion, ShouldEqual, "xxx-cros-version")
		So(resp.FaftVersion, ShouldEqual, "xxx-faft-version")
		So(resp.FirmwareVersion, ShouldEqual, "xxx-firmware-version")
		So(resp.ServoCrosVersion, ShouldEqual, "xxx-labstation-cros-version")
	})

	Convey("Test GetStableVersion RPC -- look up labstation proper", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()

		// use a fake labstation
		duts := []*inventory.DeviceUnderTest{
			{
				Common: &inventory.CommonDeviceSpecs{
					Attributes: []*inventory.KeyValue{
						{
							Key:   strptr("servo_host"),
							Value: strptr("xxx-labstation"),
						},
					},
					Id:       strptr("xxx-id"),
					Hostname: strptr("xxx-hostname"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-model"),
						Board: strptr("xxx-build-target"),
					},
				},
			},
			{
				Common: &inventory.CommonDeviceSpecs{
					Id:       strptr("xxx-labstation-id"),
					Hostname: strptr("xxx-labstation"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-labstation-model"),
						Board: strptr("xxx-labstation-board"),
					},
				},
			},
		}

		err := dsinventory.UpdateDUTs(ctx, duts)
		So(err, ShouldBeNil)

		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-labstation-board", "xxx-labstation-model", "xxx-labstation-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-labstation-board", "xxx-labstation-model", "xxx-labstation-firmware-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFaftStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-faft-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-firmware-version")
		So(err, ShouldBeNil)

		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				Hostname: "xxx-labstation",
			},
		)

		So(err, ShouldBeNil)
		So(resp.CrosVersion, ShouldEqual, "xxx-labstation-cros-version")
		So(resp.FaftVersion, ShouldEqual, "")
		So(resp.FirmwareVersion, ShouldEqual, "xxx-labstation-firmware-version")
		So(resp.ServoCrosVersion, ShouldEqual, "")
		So(resp.Reason, ShouldContainSubstring, "looked up non-satlab device hostname")
	})

	Convey("Test GetStableVersion RPC -- look up beaglebone proper", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()

		// use a fake beaglebone servo
		duts := []*inventory.DeviceUnderTest{
			{
				Common: &inventory.CommonDeviceSpecs{
					Attributes: []*inventory.KeyValue{
						{
							Key:   strptr("servo_host"),
							Value: strptr("xxx-beaglebone-servo"),
						},
					},
					Id:       strptr("xxx-id"),
					Hostname: strptr("xxx-hostname"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-model"),
						Board: strptr("xxx-build-target"),
					},
				},
			},
		}

		err := dsinventory.UpdateDUTs(ctx, duts)
		So(err, ShouldBeNil)

		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleCrosStableVersion(ctx, beagleboneServo, "", "xxx-beaglebone-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFaftStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-faft-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-firmware-version")
		So(err, ShouldBeNil)

		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				Hostname: "xxx-beaglebone-servo",
			},
		)

		So(err, ShouldBeNil)
		So(resp.CrosVersion, ShouldEqual, "xxx-beaglebone-cros-version")
		So(resp.FaftVersion, ShouldEqual, "")
		So(resp.FirmwareVersion, ShouldEqual, "")
		So(resp.ServoCrosVersion, ShouldEqual, "")
		So(resp.Reason, ShouldContainSubstring, "looks like beaglebone")
	})

	Convey("Test GetStableVersion RPC -- hostname with dummy_host", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()

		// use a fake labstation
		duts := []*inventory.DeviceUnderTest{
			{
				Common: &inventory.CommonDeviceSpecs{
					Attributes: []*inventory.KeyValue{
						{
							Key:   strptr("servo_host"),
							Value: strptr("dummy_host"),
						},
					},
					Id:       strptr("xxx-id"),
					Hostname: strptr("xxx-hostname"),
					Labels: &inventory.SchedulableLabels{
						Model: strptr("xxx-model"),
						Board: strptr("xxx-build-target"),
					},
				},
			},
		}

		err := dsinventory.UpdateDUTs(ctx, duts)
		So(err, ShouldBeNil)

		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleCrosStableVersion(ctx, "xxx-labstation-board", "xxx-labstation-model", "xxx-labstation-cros-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFaftStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-faft-version")
		So(err, ShouldBeNil)
		err = dssv.PutSingleFirmwareStableVersion(ctx, "xxx-build-target", "xxx-model", "xxx-firmware-version")
		So(err, ShouldBeNil)

		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				Hostname: "xxx-hostname",
			},
		)

		So(err, ShouldBeNil)
		So(resp.CrosVersion, ShouldEqual, "xxx-cros-version")
		So(resp.FaftVersion, ShouldEqual, "xxx-faft-version")
		So(resp.FirmwareVersion, ShouldEqual, "xxx-firmware-version")
		So(resp.ServoCrosVersion, ShouldEqual, "")
		So(resp.Reason, ShouldContainSubstring, "looked up non-satlab device hostname")
	})

	Convey("Test GetStableVersion RPC -- no stable versions exist", t, func() {
		ctx := testingContext()
		datastore.GetTestable(ctx)
		tf, validate := newTestFixtureWithContext(ctx, t)
		defer validate()
		resp, err := tf.Inventory.GetStableVersion(
			ctx,
			&fleet.GetStableVersionRequest{
				BuildTarget: "xxx-build-target",
				Model:       "xxx-model",
			},
		)
		So(err, ShouldNotBeNil)
		So(resp, ShouldBeNil)
	})
}

func withDutInfoCacheValidity(ctx context.Context, v time.Duration) context.Context {
	cfg := config.Get(ctx)
	cfg.Inventory.DutInfoCacheValidity = durationpb.New(v)
	return config.Use(ctx, cfg)
}

func withSplitInventory(ctx context.Context) context.Context {
	cfg := config.Get(ctx)
	cfg.Inventory.Multifile = true
	return config.Use(ctx, cfg)
}

// Maximum time to failure: (2^7 - 1)*(50/1000) = 6.35 seconds
var testRetriesTemplate = retry.ExponentialBackoff{
	Limited: retry.Limited{
		Delay:   50 * time.Millisecond,
		Retries: 7,
	},
	MaxDelay:   5 * time.Second,
	Multiplier: 2,
}

func testRetryIteratorFactory() retry.Iterator {
	it := testRetriesTemplate
	return &it
}

func strptr(x string) *string {
	return &x
}
