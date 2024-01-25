// Copyright 2024 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package main is the entrypoint to the fleet cost server.
package main

import (
	"net/http"
	"strings"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/config/server/cfgmodule"
	"go.chromium.org/luci/grpc/prpc"
	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/gaeemulation"
	"go.chromium.org/luci/server/module"

	"infra/cros/fleetcost/internal/costserver"
	ufspb "infra/unifiedfleet/api/v1/rpc"
)

// getUFSName gets the name of the UFS service corresponding to the current cloud project if there is one.
func getUFSName(cloudProject string) string {
	if strings.HasSuffix(cloudProject, "prod") {
		return "ufs.api.cr.dev"
	}
	return "staging.ufs.api.cr.dev"
}

// main starts the fleet cost server.
func main() {
	mods := []module.Module{
		gaeemulation.NewModuleFromFlags(),
		cfgmodule.NewModuleFromFlags(),
	}

	server.Main(nil, mods, func(srv *server.Server) error {
		ufsHostname := getUFSName(srv.Options.CloudProject)
		t, err := auth.GetRPCTransport(srv.Context, auth.AsSelf, auth.WithScopes(auth.CloudOAuthScopes...))
		if err != nil {
			return errors.Annotate(err, "setting up UFS client").Err()
		}
		httpClient := &http.Client{
			Transport: t,
		}
		prpcClient := &prpc.Client{
			C:    httpClient,
			Host: ufsHostname,
		}
		ufsClient := ufspb.NewFleetPRPCClient(prpcClient)
		fleetCostFrontend := costserver.NewFleetCostFrontend().(*costserver.FleetCostFrontend)
		costserver.SetUFSClient(fleetCostFrontend, ufsClient)
		costserver.SetUFSHostname(fleetCostFrontend, ufsHostname)
		costserver.InstallServices(fleetCostFrontend, srv)
		logging.Infof(srv.Context, "Initialization finished.")
		return nil
	})
}
