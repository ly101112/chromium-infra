// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common_commands

import (
	"context"
	"fmt"
	"infra/cros/cmd/common_lib/common"
	"infra/cros/cmd/common_lib/interfaces"
	"infra/cros/cmd/cros_test_runner/data"

	labapi "go.chromium.org/chromiumos/config/go/test/lab/api"
	"go.chromium.org/chromiumos/infra/proto/go/test_platform/skylab_test_runner"
	"go.chromium.org/luci/common/errors"
)

// ContainerStartCmd represents gcloud auth cmd.
type ContainerStartCmd struct {
	*interfaces.SingleCmdByExecutor

	// Deps
	ContainerRequest *skylab_test_runner.ContainerRequest
	ContainerImage   string

	// Updates
	Endpoint          *labapi.IpEndpoint
	ContainerInstance interfaces.ContainerInterface
}

// ExtractDependencies extracts all the command dependencies from state keeper.
func (cmd *ContainerStartCmd) ExtractDependencies(ctx context.Context,
	ski interfaces.StateKeeperInterface) error {
	var err error
	switch sk := ski.(type) {
	case *data.HwTestStateKeeper:
		err = cmd.extractDepsFromHwTestStateKeeper(ctx, sk)
	case *data.LocalTestStateKeeper:
		err = cmd.extractDepsFromHwTestStateKeeper(ctx, &sk.HwTestStateKeeper)
	default:
		return fmt.Errorf("StateKeeper '%T' is not supported by cmd type %s.", sk, cmd.GetCommandType())
	}

	if err != nil {
		return errors.Annotate(err, "error during extracting dependencies for command %s: ", cmd.GetCommandType()).Err()
	}

	return nil
}

// UpdateStateKeeper updates the state keeper with info from the cmd.
func (cmd *ContainerStartCmd) UpdateStateKeeper(
	ctx context.Context,
	ski interfaces.StateKeeperInterface) error {

	var err error
	switch sk := ski.(type) {
	case *data.HwTestStateKeeper:
		err = cmd.updateHwTestStateKeeper(ctx, sk)
	case *data.LocalTestStateKeeper:
		err = cmd.updateHwTestStateKeeper(ctx, &sk.HwTestStateKeeper)
	}

	if err != nil {
		return errors.Annotate(err, "error during updating for command %s: ", cmd.GetCommandType()).Err()
	}

	return nil
}

func (cmd *ContainerStartCmd) extractDepsFromHwTestStateKeeper(
	ctx context.Context,
	sk *data.HwTestStateKeeper) error {

	if sk.ContainerQueue.Len() < 1 {
		return fmt.Errorf("cmd %q missing dependency: ContainerRequest", cmd.GetCommandType())
	}
	cmd.ContainerRequest = sk.ContainerQueue.Remove(sk.ContainerQueue.Front()).(*skylab_test_runner.ContainerRequest)

	for _, dep := range cmd.ContainerRequest.DynamicDeps {
		if err := common.Inject(cmd.ContainerRequest.Container, dep.Key, sk.Injectables, dep.Value); err != nil {
			return fmt.Errorf("cmd %q failed injecting %s into %s, err: %s", cmd.GetCommandType(), dep.Value, dep.Key, err)
		}
	}

	containerImage, err := common.GetContainerImageFromMap(cmd.ContainerRequest.ContainerImageKey, sk.ContainerImages)
	if err != nil {
		return fmt.Errorf("cmd %q missing dependency: ContainerImage", cmd.GetCommandType())
	}
	cmd.ContainerImage = containerImage

	return nil
}

func (cmd *ContainerStartCmd) updateHwTestStateKeeper(
	ctx context.Context,
	sk *data.HwTestStateKeeper) error {

	if cmd.Endpoint != nil && cmd.ContainerRequest.DynamicIdentifier != "" {
		sk.Injectables[cmd.ContainerRequest.DynamicIdentifier] = common.ProtoToInterfaceMap(cmd.Endpoint)
	}

	if cmd.ContainerInstance != nil && cmd.ContainerRequest.DynamicIdentifier != "" {
		sk.ContainerInstances[cmd.ContainerRequest.ContainerImageKey] = cmd.ContainerInstance
	}

	return nil
}

func NewContainerStartCmd(executor interfaces.ExecutorInterface) *ContainerStartCmd {
	singleCmdByExec := interfaces.NewSingleCmdByExecutor(ContainerStartCmdType, executor)
	cmd := &ContainerStartCmd{SingleCmdByExecutor: singleCmdByExec}
	cmd.ConcreteCmd = cmd
	return cmd
}