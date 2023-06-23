// Copyright 2020 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package ufs

import (
	"context"

	"infra/cros/dutstate"

	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"google.golang.org/genproto/protobuf/field_mask"

	ufslab "infra/unifiedfleet/api/v1/models/chromeos/lab"
	ufsAPI "infra/unifiedfleet/api/v1/rpc"
)

// Allowlist of DUT states that are safe to overwrite.
var dutStatesSafeForOverwrite = map[dutstate.State]bool{
	dutstate.NeedsRepair: true,
	dutstate.Ready:       true,
	dutstate.Reserved:    true,
}

// SafeUpdateUFSDUTState attempts to safely update the DUT state to the
// given value in UFS. States other than Ready and NeedsRepair are
// ignored.
func SafeUpdateUFSDUTState(ctx context.Context, authFlags *authcli.Flags, dutName, dutState, ufsService string, repairRequests []string) error {
	c, err := NewClient(ctx, ufsService, authFlags)
	if err != nil {
		return errors.Annotate(err, "save update ufs state").Err()
	}
	info := dutstate.Read(ctx, c, dutName)
	logging.Infof(ctx, "Received DUT state from UFS: %s", info.State)
	if info.DeviceId == "" {
		return errors.Reason("save update ufs state: deviceId not found").Err()
	}
	if dutStatesSafeForOverwrite[info.State] {
		req := &ufsAPI.UpdateTestDataRequest{
			DeviceId:      info.DeviceId,
			Hostname:      dutName,
			ResourceState: dutstate.ConvertToUFSState(dutstate.State(dutState)),
		}
		maskPaths := []string{"dut.state"}
		// ReapirRequests are supported only for ChromeOS devices.
		if info.DeviceType == "chromeos" {
			// Convert repair-requests to UFS enum.
			var ufsRepairRequests []ufslab.DutState_RepairRequest
			for _, rr := range repairRequests {
				if v, ok := ufslab.DutState_RepairRequest_value[rr]; ok {
					ufsRepairRequests = append(ufsRepairRequests, ufslab.DutState_RepairRequest(v))
				} else {
					logging.Debugf(ctx, "Repair-request %q is incorrect and skipped!", rr)
				}
			}
			if len(ufsRepairRequests) > 0 {
				maskPaths = append(maskPaths, "dut_state.repair_requests")
				req.DeviceData = &ufsAPI.UpdateTestDataRequest_ChromeosData{
					ChromeosData: &ufsAPI.UpdateTestDataRequest_ChromeOs{
						DutState: &ufslab.DutState{
							RepairRequests: ufsRepairRequests,
						},
					},
				}
			}
		}
		req.UpdateMask = &field_mask.FieldMask{Paths: maskPaths}
		_, err = c.UpdateTestData(ctx, req)
		return errors.Annotate(err, "save update ufs state").Err()
	}
	logging.Warningf(ctx, "Not saving requested DUT state %s, since current DUT state is %s, which should never be overwritten!", dutState, info.State)
	return nil
}
