// Code generated by svcdec; DO NOT EDIT.

package fleet

import (
	"context"

	proto "github.com/golang/protobuf/proto"
)

type DecoratedInventory struct {
	// Service is the service to decorate.
	Service InventoryServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(c context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(c context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedInventory) DeployDut(c context.Context, req *DeployDutRequest) (rsp *DeployDutResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "DeployDut", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeployDut(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "DeployDut", rsp, err)
	}
	return
}

func (s *DecoratedInventory) RedeployDut(c context.Context, req *RedeployDutRequest) (rsp *RedeployDutResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "RedeployDut", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.RedeployDut(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "RedeployDut", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetDeploymentStatus(c context.Context, req *GetDeploymentStatusRequest) (rsp *GetDeploymentStatusResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "GetDeploymentStatus", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDeploymentStatus(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "GetDeploymentStatus", rsp, err)
	}
	return
}

func (s *DecoratedInventory) DeleteDuts(c context.Context, req *DeleteDutsRequest) (rsp *DeleteDutsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "DeleteDuts", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeleteDuts(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "DeleteDuts", rsp, err)
	}
	return
}

func (s *DecoratedInventory) EnsurePoolHealthy(c context.Context, req *EnsurePoolHealthyRequest) (rsp *EnsurePoolHealthyResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "EnsurePoolHealthy", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.EnsurePoolHealthy(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "EnsurePoolHealthy", rsp, err)
	}
	return
}

func (s *DecoratedInventory) EnsurePoolHealthyForAllModels(c context.Context, req *EnsurePoolHealthyForAllModelsRequest) (rsp *EnsurePoolHealthyForAllModelsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "EnsurePoolHealthyForAllModels", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.EnsurePoolHealthyForAllModels(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "EnsurePoolHealthyForAllModels", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ResizePool(c context.Context, req *ResizePoolRequest) (rsp *ResizePoolResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "ResizePool", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ResizePool(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ResizePool", rsp, err)
	}
	return
}

func (s *DecoratedInventory) RemoveDutsFromDrones(c context.Context, req *RemoveDutsFromDronesRequest) (rsp *RemoveDutsFromDronesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "RemoveDutsFromDrones", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.RemoveDutsFromDrones(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "RemoveDutsFromDrones", rsp, err)
	}
	return
}

func (s *DecoratedInventory) AssignDutsToDrones(c context.Context, req *AssignDutsToDronesRequest) (rsp *AssignDutsToDronesResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "AssignDutsToDrones", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.AssignDutsToDrones(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "AssignDutsToDrones", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ListServers(c context.Context, req *ListServersRequest) (rsp *ListServersResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "ListServers", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListServers(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListServers", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetDutInfo(c context.Context, req *GetDutInfoRequest) (rsp *GetDutInfoResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "GetDutInfo", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDutInfo(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "GetDutInfo", rsp, err)
	}
	return
}

func (s *DecoratedInventory) GetDroneConfig(c context.Context, req *GetDroneConfigRequest) (rsp *GetDroneConfigResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "GetDroneConfig", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetDroneConfig(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "GetDroneConfig", rsp, err)
	}
	return
}

func (s *DecoratedInventory) ListRemovedDuts(c context.Context, req *ListRemovedDutsRequest) (rsp *ListRemovedDutsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "ListRemovedDuts", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.ListRemovedDuts(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "ListRemovedDuts", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateDutLabels(c context.Context, req *UpdateDutLabelsRequest) (rsp *UpdateDutLabelsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "UpdateDutLabels", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateDutLabels(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdateDutLabels", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateCachedInventory(c context.Context, req *UpdateCachedInventoryRequest) (rsp *UpdateCachedInventoryResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "UpdateCachedInventory", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateCachedInventory(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdateCachedInventory", rsp, err)
	}
	return
}

func (s *DecoratedInventory) UpdateDeviceConfig(c context.Context, req *UpdateDeviceConfigRequest) (rsp *UpdateDeviceConfigResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(c, "UpdateDeviceConfig", req)
		if err == nil {
			c = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateDeviceConfig(c, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(c, "UpdateDeviceConfig", rsp, err)
	}
	return
}
