// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"go.chromium.org/luci/grpc/grpcutil"
	"golang.org/x/net/context"

	api "infra/appengine/unified-fleet/api/v1"
	"infra/libs/fleet/datastore"
	fleet "infra/libs/fleet/protos/go"
	"infra/libs/fleet/registration"
)

// RegistrationServerImpl implements fleet interfaces.
type RegistrationServerImpl struct {
}

// CreateMachines creates machines in datastore
func (fs *RegistrationServerImpl) CreateMachines(ctx context.Context, req *api.MachineList) (response *api.MachineResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res, err := registration.CreateMachines(ctx, req.Machine)
	if err != nil {
		return nil, err
	}
	return &api.MachineResponse{
		Passed: toMachineResult(res.Passed()),
		Failed: toMachineResult(res.Failed()),
	}, err
}

// GetMachines gets the machines information from datastore.
func (fs *RegistrationServerImpl) GetMachines(ctx context.Context, req *api.EntityIDList) (response *api.MachineResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	res := registration.GetMachinesByID(ctx, req.Id)
	return &api.MachineResponse{
		Passed: toMachineResult(res.Passed()),
		Failed: toMachineResult(res.Failed()),
	}, err
}

// ListMachines gets all the machines information from datastore.
func (fs *RegistrationServerImpl) ListMachines(ctx context.Context, req *api.ListMachinesRequest) (response *api.MachineResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	res, err := registration.GetAllMachines(ctx)
	if err != nil {
		return nil, err
	}
	return &api.MachineResponse{
		Passed: toMachineResult(res.Passed()),
		Failed: toMachineResult(res.Failed()),
	}, err
}

// UpdateMachines updates the machines information in datastore.
func (fs *RegistrationServerImpl) UpdateMachines(ctx context.Context, req *api.MachineList) (response *api.MachineResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	if err := req.Validate(); err != nil {
		return nil, err
	}
	res, err := registration.UpdateMachines(ctx, req.Machine)
	if err != nil {
		return nil, err
	}
	return &api.MachineResponse{
		Passed: toMachineResult(res.Passed()),
		Failed: toMachineResult(res.Failed()),
	}, err
}

// DeleteMachines deletes the machines from datastore.
func (fs *RegistrationServerImpl) DeleteMachines(ctx context.Context, req *api.EntityIDList) (response *api.EntityIDResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	res := registration.DeleteMachines(ctx, req.Id)
	return &api.EntityIDResponse{
		Passed: toEntityIDResult(res.Passed()),
		Failed: toEntityIDResult(res.Failed()),
	}, err
}

// CreateRacks creates Racks in datastore
func (fs *RegistrationServerImpl) CreateRacks(ctx context.Context, req *api.RackList) (response *api.RackResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	return &api.RackResponse{}, err
}

// GetRacks gets the Racks information from datastore.
func (fs *RegistrationServerImpl) GetRacks(ctx context.Context, req *api.EntityIDList) (response *api.RackResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	return &api.RackResponse{}, err
}

// ListRacks gets all the Racks information from datastore.
func (fs *RegistrationServerImpl) ListRacks(ctx context.Context, req *api.ListRacksRequest) (response *api.RackResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	return &api.RackResponse{}, err
}

// UpdateRacks updates the Racks information in datastore.
func (fs *RegistrationServerImpl) UpdateRacks(ctx context.Context, req *api.RackList) (response *api.RackResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	return &api.RackResponse{}, err
}

// DeleteRacks deletes the Racks from datastore.
func (fs *RegistrationServerImpl) DeleteRacks(ctx context.Context, req *api.EntityIDList) (response *api.EntityIDResponse, err error) {
	defer func() {
		err = grpcutil.GRPCifyAndLogErr(ctx, err)
	}()
	return &api.EntityIDResponse{}, err
}

func toMachineResult(res datastore.OpResults) []*api.MachineResult {
	cpRes := make([]*api.MachineResult, len(res))
	for i, r := range res {
		errMsg := ""
		machine := &fleet.Machine{}
		if r.Err != nil {
			errMsg = r.Err.Error()
		}
		if r.Data != nil {
			machine = r.Data.(*fleet.Machine)
		}
		cpRes[i] = &api.MachineResult{
			Machine:  machine,
			ErrorMsg: errMsg,
		}
	}
	return cpRes
}

func toEntityIDResult(res datastore.OpResults) []*api.EntityIDResult {
	cpRes := make([]*api.EntityIDResult, len(res))
	for i, r := range res {
		errMsg := ""
		machine := &fleet.Machine{}
		if r.Err != nil {
			errMsg = r.Err.Error()
		}
		if r.Data != nil {
			machine = r.Data.(*fleet.Machine)
		}
		cpRes[i] = &api.EntityIDResult{
			Id:       machine.Id.GetValue(),
			ErrorMsg: errMsg,
		}
	}
	return cpRes
}
