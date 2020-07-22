// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package controller

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"

	ufspb "infra/unifiedfleet/api/v1/proto"
	. "infra/unifiedfleet/app/model/datastore"
	"infra/unifiedfleet/app/model/registration"
)

func mockSwitch(id string) *ufspb.Switch {
	return &ufspb.Switch{
		Name: id,
	}
}

func TestCreateSwitch(t *testing.T) {
	t.Parallel()
	ctx := testingContext()
	rack1 := &ufspb.Rack{
		Name: "rack-1",
		Rack: &ufspb.Rack_ChromeBrowserRack{
			ChromeBrowserRack: &ufspb.ChromeBrowserRack{},
		},
	}
	registration.CreateRack(ctx, rack1)
	Convey("CreateSwitch", t, func() {
		Convey("Create new switch with already existing switch - error", func() {
			switch1 := &ufspb.Switch{
				Name: "switch-1",
			}
			_, err := registration.CreateSwitch(ctx, switch1)

			resp, err := CreateSwitch(ctx, switch1, "rack-5")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "Switch switch-1 already exists in the system")
		})

		Convey("Create new switch with non existing rack", func() {
			switch2 := &ufspb.Switch{
				Name: "switch-2",
			}
			resp, err := CreateSwitch(ctx, switch2, "rack-5")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "There is no Rack with RackID rack-5 in the system.")
		})

		Convey("Create new switch with existing rack with switches", func() {
			rack := &ufspb.Rack{
				Name: "rack-10",
				Rack: &ufspb.Rack_ChromeBrowserRack{
					ChromeBrowserRack: &ufspb.ChromeBrowserRack{
						Switches: []string{"switch-5"},
					},
				},
			}
			_, err := registration.CreateRack(ctx, rack)
			So(err, ShouldBeNil)

			switch1 := &ufspb.Switch{
				Name: "switch-20",
			}
			resp, err := CreateSwitch(ctx, switch1, "rack-10")
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, switch1)

			mresp, err := registration.GetRack(ctx, "rack-10")
			So(err, ShouldBeNil)
			So(mresp.GetChromeBrowserRack().GetSwitches(), ShouldResemble, []string{"switch-5", "switch-20"})
		})

		Convey("Create new switch with existing rack without switches", func() {
			rack := &ufspb.Rack{
				Name: "rack-15",
				Rack: &ufspb.Rack_ChromeBrowserRack{
					ChromeBrowserRack: &ufspb.ChromeBrowserRack{},
				},
			}
			_, err := registration.CreateRack(ctx, rack)
			So(err, ShouldBeNil)

			switch1 := &ufspb.Switch{
				Name: "switch-25",
			}
			resp, err := CreateSwitch(ctx, switch1, "rack-15")
			So(err, ShouldBeNil)
			So(resp, ShouldResembleProto, switch1)

			mresp, err := registration.GetRack(ctx, "rack-15")
			So(err, ShouldBeNil)
			So(mresp.GetChromeBrowserRack().GetSwitches(), ShouldResemble, []string{"switch-25"})
		})
	})
}

func TestUpdateSwitch(t *testing.T) {
	t.Parallel()
	ctx := testingContext()
	Convey("UpdateSwitch", t, func() {
		Convey("Update switch with non-existing switch", func() {
			rack1 := &ufspb.Rack{
				Name: "rack-1",
			}
			_, err := registration.CreateRack(ctx, rack1)
			So(err, ShouldBeNil)

			switch1 := &ufspb.Switch{
				Name: "switch-1",
			}
			resp, err := UpdateSwitch(ctx, switch1, "rack-1")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "There is no Switch with SwitchID switch-1 in the system")
		})

		Convey("Update switch with new rack", func() {
			rack3 := &ufspb.Rack{
				Name: "rack-3",
				Rack: &ufspb.Rack_ChromeBrowserRack{
					ChromeBrowserRack: &ufspb.ChromeBrowserRack{
						Switches: []string{"switch-3"},
					},
				},
			}
			_, err := registration.CreateRack(ctx, rack3)
			So(err, ShouldBeNil)

			rack4 := &ufspb.Rack{
				Name: "rack-4",
				Rack: &ufspb.Rack_ChromeBrowserRack{
					ChromeBrowserRack: &ufspb.ChromeBrowserRack{
						Switches: []string{"switch-4"},
					},
				},
			}
			_, err = registration.CreateRack(ctx, rack4)
			So(err, ShouldBeNil)

			switch3 := &ufspb.Switch{
				Name: "switch-3",
			}
			_, err = registration.CreateSwitch(ctx, switch3)
			So(err, ShouldBeNil)

			resp, err := UpdateSwitch(ctx, switch3, "rack-4")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(resp, ShouldResembleProto, switch3)

			mresp, err := registration.GetRack(ctx, "rack-3")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(mresp.GetChromeBrowserRack().GetSwitches(), ShouldBeNil)

			mresp, err = registration.GetRack(ctx, "rack-4")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(mresp.GetChromeBrowserRack().GetSwitches(), ShouldResemble, []string{"switch-4", "switch-3"})
		})

		Convey("Update switch with same rack", func() {
			rack := &ufspb.Rack{
				Name: "rack-5",
				Rack: &ufspb.Rack_ChromeBrowserRack{
					ChromeBrowserRack: &ufspb.ChromeBrowserRack{
						Switches: []string{"switch-5"},
					},
				},
			}
			_, err := registration.CreateRack(ctx, rack)
			So(err, ShouldBeNil)

			switch1 := &ufspb.Switch{
				Name: "switch-5",
			}
			_, err = registration.CreateSwitch(ctx, switch1)
			So(err, ShouldBeNil)

			resp, err := UpdateSwitch(ctx, switch1, "rack-5")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
			So(resp, ShouldResembleProto, switch1)
		})

		Convey("Update switch with non existing rack", func() {
			switch1 := &ufspb.Switch{
				Name: "switch-6",
			}
			_, err := registration.CreateSwitch(ctx, switch1)
			So(err, ShouldBeNil)

			resp, err := UpdateSwitch(ctx, switch1, "rack-6")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "There is no Rack with RackID rack-6 in the system.")
		})

	})
}

func TestDeleteSwitch(t *testing.T) {
	t.Parallel()
	ctx := testingContext()
	switch1 := mockSwitch("switch-1")
	switch2 := mockSwitch("switch-2")
	Convey("DeleteSwitch", t, func() {
		Convey("Delete switch by existing ID with machine reference", func() {
			resp, cerr := registration.CreateSwitch(ctx, switch1)
			So(cerr, ShouldBeNil)
			So(resp, ShouldResembleProto, switch1)

			nic := &ufspb.Nic{
				Name: "machine1-eth0",
				SwitchInterface: &ufspb.SwitchInterface{
					Switch: "switch-1",
				},
			}
			mresp, merr := registration.CreateNic(ctx, nic)
			So(merr, ShouldBeNil)
			So(mresp, ShouldResembleProto, nic)

			err := DeleteSwitch(ctx, "switch-1")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, CannotDelete)

			resp, cerr = registration.GetSwitch(ctx, "switch-1")
			So(resp, ShouldNotBeNil)
			So(cerr, ShouldBeNil)
			So(resp, ShouldResembleProto, switch1)
		})
		Convey("Delete switch successfully by existing ID without references", func() {
			resp, cerr := registration.CreateSwitch(ctx, switch2)
			So(cerr, ShouldBeNil)
			So(resp, ShouldResembleProto, switch2)

			err := DeleteSwitch(ctx, "switch-2")
			So(err, ShouldBeNil)

			resp, cerr = registration.GetSwitch(ctx, "switch-2")
			So(resp, ShouldBeNil)
			So(cerr, ShouldNotBeNil)
			So(cerr.Error(), ShouldContainSubstring, NotFound)
		})
	})
}
