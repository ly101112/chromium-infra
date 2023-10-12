// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package vm

import (
	"fmt"
	"os"

	"github.com/maruel/subcommands"
	"go.chromium.org/luci/common/cli"

	"infra/cmdsupport/cmdlib"
	"infra/vm_leaser/client"
)

const leasesCmd = "leases"

var leases = &subcommands.Command{
	UsageLine: fmt.Sprintf("%s [FLAGS...]", leasesCmd),
	ShortDesc: "Print a list of the current user's leases",
	LongDesc: `Print a list of the current user's leases.

This command's behavior is subject to change without notice.
Do not build automation around this subcommand.`,
	CommandRun: func() subcommands.CommandRun {
		c := &leasesRun{}
		return c
	},
}

type leasesRun struct {
	subcommands.CommandRunBase
}

func (c *leasesRun) Run(a subcommands.Application, _ []string, env subcommands.Env) int {
	if err := c.innerRun(a, env); err != nil {
		cmdlib.PrintError(a, err)
		return 1
	}
	return 0
}

func (c *leasesRun) innerRun(a subcommands.Application, env subcommands.Env) error {
	ctx := cli.GetContext(a, c, env)

	vmLeaser, err := client.NewClient(ctx, client.LocalConfig())
	if err != nil {
		return err
	}

	vms, err := listLeases(vmLeaser, ctx)
	if err != nil {
		return err
	}

	if len(vms) == 0 {
		fmt.Println("No active VM leases")
		return nil
	}

	fmt.Printf("%d active lease(s)\n", len(vms))
	printVMList(vms, os.Stdout)

	return nil
}
