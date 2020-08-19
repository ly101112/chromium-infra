// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package rpm

import (
	"fmt"

	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/grpc/prpc"

	"infra/cmd/shivas/site"
	"infra/cmd/shivas/utils"
	"infra/cmdsupport/cmdlib"
	ufspb "infra/unifiedfleet/api/v1/proto"
	ufsAPI "infra/unifiedfleet/api/v1/rpc"
	ufsUtil "infra/unifiedfleet/app/util"
)

// GetRPMCmd get rpm by given name.
var GetRPMCmd = &subcommands.Command{
	UsageLine: "rpm {RPM Name}",
	ShortDesc: "Get rpm details by name",
	LongDesc: `Get rpm details by name.

Example:

shivas get rpm {RPM Name}
Gets the rpm and prints the output in JSON format.`,
	CommandRun: func() subcommands.CommandRun {
		c := &getRPM{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.envFlags.Register(&c.Flags)
		c.commonFlags.Register(&c.Flags)
		c.outputFlags.Register(&c.Flags)
		return c
	},
}

type getRPM struct {
	subcommands.CommandRunBase
	authFlags   authcli.Flags
	envFlags    site.EnvFlags
	commonFlags site.CommonFlags
	outputFlags site.OutputFlags
}

func (c *getRPM) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.innerRun(a, args, env); err != nil {
		cmdlib.PrintError(a, err)
		return 1
	}
	return 0
}

func (c *getRPM) innerRun(a subcommands.Application, args []string, env subcommands.Env) error {
	if err := c.validateArgs(); err != nil {
		return err
	}
	ctx := cli.GetContext(a, c, env)
	ctx = utils.SetupContext(ctx)
	hc, err := cmdlib.NewHTTPClient(ctx, &c.authFlags)
	if err != nil {
		return err
	}
	e := c.envFlags.Env()
	if c.commonFlags.Verbose() {
		fmt.Printf("Using UnifiedFleet service %s\n", e.UnifiedFleetService)
	}
	ic := ufsAPI.NewFleetPRPCClient(&prpc.Client{
		C:       hc,
		Host:    e.UnifiedFleetService,
		Options: site.DefaultPRPCOptions,
	})
	res, err := ic.GetRPM(ctx, &ufsAPI.GetRPMRequest{
		Name: ufsUtil.AddPrefix(ufsUtil.RPMCollection, args[0]),
	})
	if err != nil {
		return err
	}
	res.Name = ufsUtil.RemovePrefix(res.Name)
	return c.print(res)
}

func (c *getRPM) print(rpm *ufspb.RPM) error {
	if c.outputFlags.JSON() {
		utils.PrintProtoJSON(rpm, c.outputFlags.Emit())
		return nil
	}
	if c.outputFlags.Tsv() {
		utils.PrintTSVRPMs([]*ufspb.RPM{rpm}, false)
		return nil
	}
	utils.PrintTitle(utils.RpmTitle)
	utils.PrintRPMs([]*ufspb.RPM{rpm}, false)
	return nil
}

func (c *getRPM) validateArgs() error {
	if c.Flags.NArg() == 0 {
		return cmdlib.NewUsageError(c.Flags, "Please provide the rpm name.")
	}
	return nil
}
