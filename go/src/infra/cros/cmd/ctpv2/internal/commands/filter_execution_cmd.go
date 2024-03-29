// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package commands

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	testapi "go.chromium.org/chromiumos/config/go/test/api"
	"go.chromium.org/luci/common/errors"

	"infra/cros/cmd/common_lib/common"
	"infra/cros/cmd/common_lib/interfaces"
	"infra/cros/cmd/ctpv2/data"
)

// FilterExecutionCmd represents test execution cmd.
type FilterExecutionCmd struct {
	*interfaces.SingleCmdByExecutor

	// Deps
	InputTestPlan *testapi.InternalTestplan
	ContainerInfo *data.ContainerInfo

	// Updates
	OutputTestPlan *testapi.InternalTestplan
}

// ExtractDependencies extracts all the command dependencies from state keeper.
func (cmd *FilterExecutionCmd) ExtractDependencies(
	ctx context.Context,
	ski interfaces.StateKeeperInterface) error {

	var err error
	switch sk := ski.(type) {
	case *data.FilterStateKeeper:
		err = cmd.extractDepsFromFilterStateKeeper(ctx, sk)

	default:
		return fmt.Errorf("StateKeeper '%T' is not supported by cmd type %s.", sk, cmd.GetCommandType())
	}

	if err != nil {
		return errors.Annotate(err, "error during extracting dependencies for command %s: ", cmd.GetCommandType()).Err()
	}

	return nil
}

// UpdateStateKeeper updates the state keeper with info from the cmd.
func (cmd *FilterExecutionCmd) UpdateStateKeeper(
	ctx context.Context,
	ski interfaces.StateKeeperInterface) error {

	var err error
	switch sk := ski.(type) {
	case *data.FilterStateKeeper:
		err = cmd.updateFilterStateKeeper(ctx, sk)
	}

	if err != nil {
		return errors.Annotate(err, "error during updating for command %s: ", cmd.GetCommandType()).Err()
	}

	return nil
}

func (cmd *FilterExecutionCmd) extractDepsFromFilterStateKeeper(
	ctx context.Context,
	sk *data.FilterStateKeeper) error {

	if sk.ContainerInfoQueue.Len() < 1 {
		return fmt.Errorf("cmd %q missing dependency: ContainerInfo", cmd.GetCommandType())
	}

	cmd.ContainerInfo = sk.ContainerInfoQueue.Remove(sk.ContainerInfoQueue.Front()).(*data.ContainerInfo)
	if cmd.ContainerInfo.ServiceEndpoint == nil {
		return fmt.Errorf("cmd %q missing dependency: ServiceEndpoint", cmd.GetCommandType())
	}

	if sk.TestPlanStates == nil || len(sk.TestPlanStates) == 0 {
		if sk.InitialInternalTestPlan != nil {
			// Set the first state from initial test plan
			sk.TestPlanStates = append(sk.TestPlanStates, sk.InitialInternalTestPlan)
			// Set the cmd input test plan
			cmd.InputTestPlan = proto.Clone(sk.InitialInternalTestPlan).(*testapi.InternalTestplan)
		} else {
			return fmt.Errorf("Cmd %q missing dependency: InputTestPlan", cmd.GetCommandType())
		}
	} else {
		// Get the last test plan state and set it as input test plan for current filter
		cmd.InputTestPlan = proto.Clone(sk.TestPlanStates[len(sk.TestPlanStates)-1]).(*testapi.InternalTestplan)
	}

	// TODO (azrahman): remove these custom test plans call once ttcp filter stablized.
	// Only to be used to test ttcp filter through led.
	// if cmd.ContainerInfo.GetKey() == "ttcp-demo" {
	// 	cmd.InputTestPlan = addNTests(5)
	// }

	// if cmd.ContainerInfo.GetKey() == "ttcp-demo" {
	// 	cmd.InputTestPlan = addCustomTests()
	// }

	return nil
}

func (cmd *FilterExecutionCmd) updateFilterStateKeeper(ctx context.Context, sk *data.FilterStateKeeper) error {
	if err := common.ValidateTestPlans(cmd.InputTestPlan, cmd.OutputTestPlan); err != nil {
		return fmt.Errorf("Cmd %q failed with test plan validation: %s", cmd.GetCommandType(), err)
	}

	// Add the validated output testplan to test plan states.
	sk.TestPlanStates = append(sk.TestPlanStates, cmd.OutputTestPlan)

	return nil
}

func NewFilterExecutionCmd(executor interfaces.ExecutorInterface) *FilterExecutionCmd {
	singleCmdByExec := interfaces.NewSingleCmdByExecutor(FilterExecutionCmdType, executor)
	cmd := &FilterExecutionCmd{SingleCmdByExecutor: singleCmdByExec}
	cmd.ConcreteCmd = cmd
	return cmd
}
