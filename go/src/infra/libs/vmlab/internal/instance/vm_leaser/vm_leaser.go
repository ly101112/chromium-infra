// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package vmleaser

import (
	"context"
	"errors"
	"fmt"

	"go.chromium.org/chromiumos/config/go/test/api"
	"google.golang.org/grpc"

	vmlabpb "infra/libs/vmlab/api"
	"infra/vm_leaser/client"
)

// vmLeaserInstanceApi implements vmlabpb.InstanceApi
//
// The struct itself doesn't need to be public.
type vmLeaserInstanceApi struct{}

// New constructs a new vmlabpb.InstanceApi with VM Leaser Service backend.
func New() (vmlabpb.InstanceApi, error) {
	return &vmLeaserInstanceApi{}, nil
}

// vmLeaserServiceClient interfaces the VM Leaser service client.
type vmLeaserServiceClient interface {
	LeaseVM(ctx context.Context, r *api.LeaseVMRequest, opts ...grpc.CallOption) (*api.LeaseVMResponse, error)
	ReleaseVM(ctx context.Context, r *api.ReleaseVMRequest, opts ...grpc.CallOption) (*api.ReleaseVMResponse, error)
}

// envConfig returns a VM Leaser client config.
//
// The appropriate config is based on the environment the CLI leasing client
// wishes to connect to.
func envConfig(backendEnv vmlabpb.Config_VmLeaserBackend_Environment) *client.Config {
	switch backendEnv {
	case vmlabpb.Config_VmLeaserBackend_ENV_LOCAL:
		return client.LocalConfig()
	case vmlabpb.Config_VmLeaserBackend_ENV_STAGING:
		return client.StagingConfig()
	default:
		return client.ProdConfig()
	}
}

// Create takes a CreateVmInstanceRequest and returns a VmInstance.
//
// Create parses the CreateVmInstanceRequest into a LeaseVMRequest. A client
// connection is established with the VM Leaser service and the request is sent.
// The connected service (local/staging/production) is based on the Env config.
func (g *vmLeaserInstanceApi) Create(ctx context.Context, req *vmlabpb.CreateVmInstanceRequest) (*vmlabpb.VmInstance, error) {
	err := req.ValidateVmLeaserBackend()
	if err != nil {
		return nil, err
	}

	vmLeaser, err := client.NewClient(ctx, envConfig(req.GetConfig().GetVmLeaserBackend().GetEnv()))
	if err != nil {
		return nil, fmt.Errorf("failed to create new client: %w", err)
	}
	defer vmLeaser.Close()

	return g.leaseVM(ctx, vmLeaser.VMLeaserClient, req)
}

func (g *vmLeaserInstanceApi) Delete(ctx context.Context, ins *vmlabpb.VmInstance) error {
	vmLeaserBackend := ins.GetConfig().GetVmLeaserBackend()
	if vmLeaserBackend == nil {
		return fmt.Errorf("invalid argument: bad backend: want vm leaser, got %v", ins.GetConfig())
	}
	if ins.GetName() == "" {
		return errors.New("instance name must be set")
	}
	if vmLeaserBackend.GetVmRequirements().GetGceProject() == "" {
		return errors.New("project must be set")
	}
	if vmLeaserBackend.GetVmRequirements().GetGceRegion() == "" {
		return errors.New("zone must be set")
	}

	vmLeaser, err := client.NewClient(ctx, envConfig(vmLeaserBackend.GetEnv()))
	if err != nil {
		return fmt.Errorf("failed to create new client: %w", err)
	}
	defer vmLeaser.Close()

	return g.releaseVM(ctx, vmLeaser.VMLeaserClient, ins)
}

func (g *vmLeaserInstanceApi) List(ctx context.Context, req *vmlabpb.ListVmInstancesRequest) ([]*vmlabpb.VmInstance, error) {
	return []*vmlabpb.VmInstance{}, errors.New("not implemented")
}

// leaseVM calls LeaseVM using the VM Leaser service client.
func (g *vmLeaserInstanceApi) leaseVM(ctx context.Context, client vmLeaserServiceClient, req *vmlabpb.CreateVmInstanceRequest) (*vmlabpb.VmInstance, error) {
	vmLeaserBackend := req.GetConfig().GetVmLeaserBackend()
	rsp, err := client.LeaseVM(ctx, &api.LeaseVMRequest{
		LeaseDuration: vmLeaserBackend.GetLeaseDuration(),
		HostReqs:      vmLeaserBackend.GetVmRequirements(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to lease VM: %w", err)
	}

	return &vmlabpb.VmInstance{
		Name: rsp.GetLeaseId(),
		Ssh: &vmlabpb.AddressPort{
			Address: rsp.GetVm().GetAddress().GetHost(),
			Port:    rsp.GetVm().GetAddress().GetPort(),
		},
		Config: req.GetConfig(),
	}, nil
}

func (g *vmLeaserInstanceApi) releaseVM(ctx context.Context, client vmLeaserServiceClient, ins *vmlabpb.VmInstance) error {
	vmLeaserBackend := ins.GetConfig().GetVmLeaserBackend()
	_, err := client.ReleaseVM(ctx, &api.ReleaseVMRequest{
		LeaseId:    ins.GetName(),
		GceProject: vmLeaserBackend.GetVmRequirements().GetGceProject(),
		GceRegion:  vmLeaserBackend.GetVmRequirements().GetGceRegion(),
	})
	if err != nil {
		return fmt.Errorf("failed to release VM: %w", err)
	}
	return nil
}
