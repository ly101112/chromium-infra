// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package machine

import (
	"fmt"
	"os"

	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/grpc/prpc"
	"infra/cmd/shivas/site"
	"infra/cmd/shivas/utils"
	"infra/cmdsupport/cmdlib"
	UfleetAPI "infra/unifiedfleet/api/v1/rpc"
	UfleetUtil "infra/unifiedfleet/app/util"
)

// DeleteMachineCmd delete Machine by given name.
var DeleteMachineCmd = &subcommands.Command{
	UsageLine: "del",
	ShortDesc: "delete Machine by name",
	LongDesc: `delete Machine by name.
	./shivas machine del {Machine Name}
	Deletes the given Machine.`,
	CommandRun: func() subcommands.CommandRun {
		c := &deleteMachine{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.envFlags.Register(&c.Flags)
		return c
	},
}

type deleteMachine struct {
	subcommands.CommandRunBase
	authFlags authcli.Flags
	envFlags  site.EnvFlags
}

func (c *deleteMachine) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.innerRun(a, args, env); err != nil {
		cmdlib.PrintError(a, err)
		return 1
	}
	return 0
}

func (c *deleteMachine) innerRun(a subcommands.Application, args []string, env subcommands.Env) error {
	if err := c.validateArgs(); err != nil {
		return err
	}
	ctx := cli.GetContext(a, c, env)
	ctx = utils.SetupContext(ctx)
	hc, err := cmdlib.NewHTTPClient(ctx, &c.authFlags)
	if err != nil {
		return err
	}
	prompt := utils.CLIPrompt(a.GetOut(), os.Stdin, false)
	if !prompt(fmt.Sprintf("Are you sure you want to delete Machine: %s", args[0])) {
		return nil
	}
	e := c.envFlags.Env()
	fmt.Printf("Using UnifiedFleet service %s\n", e.UnifiedFleetService)
	ic := UfleetAPI.NewFleetPRPCClient(&prpc.Client{
		C:       hc,
		Host:    e.UnifiedFleetService,
		Options: site.DefaultPRPCOptions,
	})
	_, err = ic.DeleteMachine(ctx, &UfleetAPI.DeleteMachineRequest{
		Name: UfleetUtil.AddPrefix(UfleetUtil.MachineCollection, args[0]),
	})
	if err == nil {
		fmt.Fprintln(a.GetOut(), args[0], "deleted successfully.")
		return nil
	}
	return err
}

func (c *deleteMachine) validateArgs() error {
	if c.Flags.NArg() == 0 {
		return cmdlib.NewUsageError(c.Flags, "Please provide a Machine Name")
	}
	return nil
}
