// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package executor

import (
	"context"
	"fmt"

	managers "infra/cros/cmd/cros_test_platformV2/docker_managers"

	testapi "go.chromium.org/chromiumos/config/go/test/api"
)

// CrosProvisionExecutor represents executor
// for all crostoolrunner (ctr) related commands.
type CtrExecutor struct {
	Ctr                        managers.ContainerManager
	CrosProvisionServiceClient testapi.GenericProvisionServiceClient
	KeyLocation                string
}

func NewCtrExecutor(ctr managers.ContainerManager) *CtrExecutor {
	return &CtrExecutor{Ctr: ctr}
}

func (ex *CtrExecutor) Execute(ctx context.Context, cmd string) error {
	if cmd == "run" {
		fmt.Println("CTR Run.")
		return nil //ex.gcloudAuthCommandExecution(ctx)
	} else if cmd == "init" {
		fmt.Println("CTR init")
		ex.Ctr.StartManager(ctx, "foo")
		return nil
	} else if cmd == "stop" {
		fmt.Println("CTR stop")
		ex.Ctr.StopManager(ctx, "foo")
		return nil
	}
	return fmt.Errorf("invalid command given: %s", cmd)
}