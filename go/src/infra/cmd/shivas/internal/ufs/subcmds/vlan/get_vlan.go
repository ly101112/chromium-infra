// Copyright 2020 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package vlan

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/maruel/subcommands"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/common/cli"
	"go.chromium.org/luci/common/flag"
	"go.chromium.org/luci/grpc/prpc"

	"infra/cmd/shivas/cmdhelp"
	"infra/cmd/shivas/site"
	"infra/cmd/shivas/utils"
	"infra/cmdsupport/cmdlib"
	ufspb "infra/unifiedfleet/api/v1/models"
	ufsAPI "infra/unifiedfleet/api/v1/rpc"
	ufsUtil "infra/unifiedfleet/app/util"
)

// GetVlanCmd get vlan by given name.
var GetVlanCmd = &subcommands.Command{
	UsageLine: "vlan ...",
	ShortDesc: "Get vlan details by filters",
	LongDesc: `Get vlan details by filters.

Example:

shivas get vlan name1 name2

shivas get vlan

shivas get vlan -state serving

Gets the vlan and prints the output in the user-specified format.`,
	CommandRun: func() subcommands.CommandRun {
		c := &getVlan{}
		c.authFlags.Register(&c.Flags, site.DefaultAuthOptions)
		c.envFlags.Register(&c.Flags)
		c.commonFlags.Register(&c.Flags)
		c.outputFlags.Register(&c.Flags)

		c.Flags.IntVar(&c.pageSize, "n", 0, cmdhelp.ListPageSizeDesc)
		c.Flags.BoolVar(&c.keysOnly, "keys", false, cmdhelp.KeysOnlyText)
		c.Flags.BoolVar(&c.listIPs, "ips", false, cmdhelp.ShowIPsText)

		c.Flags.Var(flag.StringSlice(&c.states), "state", "Name(s) of a state to filter by. Can be specified multiple times."+cmdhelp.StateFilterHelpText)
		c.Flags.Var(flag.StringSlice(&c.zones), "zone", "Name(s) of a zone to filter by. Can be specified multiple times."+cmdhelp.ZoneFilterHelpText)
		c.Flags.Var(flag.StringSlice(&c.tags), "tag", "Name(s) of a tag to filter by. Can be specified multiple times.")

		return c
	},
}

type getVlan struct {
	subcommands.CommandRunBase
	authFlags   authcli.Flags
	envFlags    site.EnvFlags
	commonFlags site.CommonFlags
	outputFlags site.OutputFlags

	// Filters
	states []string
	zones  []string
	tags   []string

	pageSize int
	keysOnly bool
	listIPs  bool
}

func (c *getVlan) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if err := c.innerRun(a, args, env); err != nil {
		cmdlib.PrintError(a, err)
		return 1
	}
	return 0
}

func (c *getVlan) innerRun(a subcommands.Application, args []string, env subcommands.Env) error {
	if err := c.validateArgs(); err != nil {
		return err
	}

	ctx := cli.GetContext(a, c, env)
	ns, err := c.envFlags.Namespace(nil, "")
	if err != nil {
		return err
	}
	ctx = utils.SetupContext(ctx, ns)
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
	emit := !utils.NoEmitMode(c.outputFlags.NoEmit())
	full := utils.FullMode(c.outputFlags.Full())
	var res []proto.Message
	if len(args) > 0 {
		res = utils.ConcurrentGet(ctx, ic, args, c.getSingle)
	} else {
		res, err = utils.BatchList(ctx, ic, listVlans, c.formatFilters(), c.pageSize, c.keysOnly, full)
	}
	if err != nil {
		return err
	}
	err = utils.PrintEntities(ctx, ic, res, utils.PrintVlansJSON, printVlanFull, printVlanNormal,
		c.outputFlags.JSON(), emit, full, c.outputFlags.Tsv(), c.keysOnly)
	if err != nil {
		return err
	}

	if len(res) == 1 && c.listIPs {
		// Only print IP utilization in json mode for 1 vlan
		if err := printUsedIPs(ctx, ic, res); err != nil {
			cmdlib.PrintError(a, err)
		}
	}
	return nil
}

func (c *getVlan) formatFilters() []string {
	filters := make([]string, 0)
	filters = utils.JoinFilters(filters, utils.PrefixFilters(ufsUtil.StateFilterName, c.states)...)
	filters = utils.JoinFilters(filters, utils.PrefixFilters(ufsUtil.ZoneFilterName, c.zones)...)
	filters = utils.JoinFilters(filters, utils.PrefixFilters(ufsUtil.TagFilterName, c.tags)...)
	return filters
}

func (c *getVlan) getSingle(ctx context.Context, ic ufsAPI.FleetClient, name string) (proto.Message, error) {
	return ic.GetVlan(ctx, &ufsAPI.GetVlanRequest{
		Name: ufsUtil.AddPrefix(ufsUtil.VlanCollection, name),
	})
}

func listVlans(ctx context.Context, ic ufsAPI.FleetClient, pageSize int32, pageToken, filter string, keysOnly, full bool) ([]proto.Message, string, error) {
	req := &ufsAPI.ListVlansRequest{
		PageSize:  pageSize,
		PageToken: pageToken,
		Filter:    filter,
		KeysOnly:  keysOnly,
	}
	res, err := ic.ListVlans(ctx, req)
	if err != nil {
		return nil, "", err
	}
	protos := make([]proto.Message, len(res.GetVlans()))
	for i, m := range res.GetVlans() {
		protos[i] = m
	}
	return protos, res.GetNextPageToken(), nil
}

func printVlanFull(ctx context.Context, ic ufsAPI.FleetClient, msgs []proto.Message, tsv bool) error {
	return printVlanNormal(msgs, tsv, false)
}

func printVlanNormal(entities []proto.Message, tsv, keysOnly bool) error {
	if len(entities) == 0 {
		return nil
	}
	if tsv {
		utils.PrintTSVVlans(entities, keysOnly)
		return nil
	}
	utils.PrintTableTitle(utils.VlanTitle, tsv, keysOnly)
	utils.PrintVlans(entities, keysOnly)
	return nil
}

func printUsedIPs(ctx context.Context, ic ufsAPI.FleetClient, res []proto.Message) error {
	v := res[0].(*ufspb.Vlan)
	toPrint := make([]*ufspb.IP, 0)
	curPageToken := ""
	for {
		resIPs, err := ic.ListIPs(ctx, &ufsAPI.ListIPsRequest{
			Filter:    utils.PrefixFilters(ufsUtil.VlanFilterName, []string{ufsUtil.RemovePrefix(v.GetName())})[0],
			PageSize:  ufsUtil.MaxPageSize,
			PageToken: curPageToken,
		})
		if err != nil {
			return err
		}
		for _, ip := range resIPs.GetIps() {
			if ip.GetOccupied() || ip.GetReserve() {
				toPrint = append(toPrint, ip)
			}
		}
		nextPageToken := resIPs.GetNextPageToken()
		if nextPageToken == "" {
			break
		} else {
			curPageToken = nextPageToken
		}
	}

	if len(toPrint) > 0 {
		fmt.Println()
		utils.PrintTableTitle(utils.IPTitle, false, false)
		utils.PrintIPs(toPrint, false)
	}
	return nil
}

func (c *getVlan) validateArgs() error {
	if c.Flags.NArg() == 0 && c.listIPs {
		return cmdlib.NewUsageError(c.Flags, "Please provide the vlan name if `-ips` is specified.")
	}
	if c.Flags.NArg() > 1 && c.listIPs {
		return cmdlib.NewUsageError(c.Flags, "Please only provide 1 vlan name if `-ips` is specified.")
	}
	if c.listIPs && c.outputFlags.JSON() {
		return cmdlib.NewUsageError(c.Flags, "`-ips` is not supported when getting a json output of vlan.")
	}
	return nil
}
