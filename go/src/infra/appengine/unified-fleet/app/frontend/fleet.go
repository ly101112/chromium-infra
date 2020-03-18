// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"go.chromium.org/luci/common/logging"
	"golang.org/x/net/context"

	api "infra/appengine/unified-fleet/api/v1"
)

// FleetServerImpl implements fleet interfaces.
type FleetServerImpl struct {
}

// RegisterMachines registers...
func (fs *FleetServerImpl) RegisterMachines(ctx context.Context, req *api.RegisterMachinesRequest) (response *api.RegisterMachinesResponse, err error) {
	logging.Debugf(ctx, "enter RegisterMachines")
	return &api.RegisterMachinesResponse{}, err
}
