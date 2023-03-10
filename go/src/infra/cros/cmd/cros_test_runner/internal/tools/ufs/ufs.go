// Copyright 2023 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package ufs provides option to build connection to UFS service & invoke it's
// endpoints.
package ufs

import (
	"context"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/grpc/prpc"

	"infra/cros/cmd/cros_test_runner/common"
	"infra/cros/dutstate"
	ufsAPI "infra/unifiedfleet/api/v1/rpc"
)

// Allowlist of DUT states that are safe to overwrite.
var dutStatesSafeForOverwrite = map[dutstate.State]bool{
	dutstate.NeedsRepair: true,
	dutstate.Ready:       true,
	dutstate.Reserved:    true,
}

// prpcOptions is used for UFS PRPC clients.
var prpcOptions = prpcOptionWithUserAgent("skylab_local_state/6.0.0")

// NewClient initialize and return new client to work with UFS service.
func NewClient(ctx context.Context) (ufsAPI.FleetClient, error) {
	authFlags := authcli.Flags{}
	authOpts, err := authFlags.Options()
	if err != nil {
		return nil, errors.Annotate(err, "create UFS client").Err()
	}
	a := auth.NewAuthenticator(ctx, auth.SilentLogin, authOpts)
	httpClient, err := a.Client()
	if err != nil {
		return nil, errors.Annotate(err, "create UFS client").Err()
	}
	ufsClient := ufsAPI.NewFleetPRPCClient(&prpc.Client{
		C:       httpClient,
		Host:    common.UfsServiceUrl,
		Options: prpcOptions,
	})
	return ufsClient, nil
}

// prpcOptionWithUserAgent create prpc option with custom UserAgent.
// DefaultOptions provides Retry ability in case we have issue with service.
func prpcOptionWithUserAgent(userAgent string) *prpc.Options {
	options := *prpc.DefaultOptions()
	options.UserAgent = userAgent
	return &options
}

// SafeUpdateUFSDUTState attempts to safely update the DUT state to the
// given value in UFS.
func SafeUpdateUFSDUTState(ctx context.Context, dutName string, dutState dutstate.State) error {
	currentDUTState, err := GetDutStateFromUFS(ctx, dutName)
	if err != nil {
		return errors.Annotate(err, "update dut state").Err()
	}
	if dutStatesSafeForOverwrite[currentDUTState] {
		logging.Infof(ctx, "Overwriting dut state...")
		return updateDUTStateToUFS(ctx, dutName, dutState)
	}
	logging.Warningf(ctx, "Not saving requested DUT state %s, since current DUT state is %s, which should never be overwritten.", dutState, currentDUTState)
	return nil
}

// updateDUTStateToUFS send DUT state to the UFS service.
func updateDUTStateToUFS(ctx context.Context, dutName string, dutState dutstate.State) error {
	ufsClient, err := NewClient(ctx)
	if err != nil {
		return errors.Annotate(err, "update dut state").Err()
	}
	err = dutstate.Update(ctx, ufsClient, dutName, dutState)
	if err != nil {
		return errors.Annotate(err, "update dut state").Err()
	}
	return nil
}

// GetDutStateFromUFS reads DUT state from the UFS service.
func GetDutStateFromUFS(ctx context.Context, dutName string) (dutstate.State, error) {
	ufsClient, err := NewClient(ctx)
	if err != nil {
		return "", errors.Annotate(err, "get dut state").Err()
	}
	info := dutstate.Read(ctx, ufsClient, dutName)
	logging.Infof(ctx, "Received DUT state from UFS: %s", info.State)
	return info.State, nil
}