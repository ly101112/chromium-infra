// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/types/known/timestamppb"

	gceproviderpb "go.chromium.org/luci/gce/api/config/v1"
)

const USAGE = `
$ ./gce_quota_checker [--project some-project-ID] path/to/vms.cfg path/to/another/vms.cfg

This will sum up all deployments in the specified vms.cfg files for the
specified GCP project, then check those against the allowed quotas for the
project. This checks:
- VM instances per region
- instanced per VPC network
- in-use IPs per region
- CPUs per region (N1)
- N2 CPUs per region
- HDD per region
- SSD per region
- local SSD per region + family

Hints:
- region == "us-east1"
- zone == "us-east1-d"
- family == "n1", "n2", "e2", etc

NOTE: need to run 'gcloud auth application-default login' locally first.
`

type quotaVals struct {
	max  int64
	used int64
}

type regionQuotas struct {
	// One-per instance quotas
	instancesQuota quotaVals
	ipsQuota       quotaVals

	// CPUs
	cpusQuota   quotaVals
	n2CpusQuota quotaVals

	// Disk
	hddQuota       quotaVals
	remoteSSDQuota quotaVals
	// GCE tracks this one across both region and VM family (n1, e2, etc).
	localSSDPerFamilyQuota map[string]*quotaVals
}

func parseFlags() (string, bool, []string) {
	gcpProject := flag.String("project", "google.com:chromecompute", "ID of the project to get quota for.")
	isVerbose := flag.Bool("verbose", false, "Prints full quota usage.")
	flag.Usage = func() {
		fmt.Printf("%v\n", USAGE)
		os.Exit(1)
	}
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		flag.Usage()
	}
	return *gcpProject, *isVerbose, paths
}

func queryTimeSeriesQuota(ctx context.Context, client *monitoring.MetricClient, quotaName string, project string) *monitoring.TimeSeriesIterator {
	// Get all points in the metric in the past 1 day. Some of the metrics
	// don't seem to get updated very often, so need a big window.
	startTime := time.Now().UTC().Add(time.Hour * -24)
	endTime := time.Now().UTC()
	monitoringRequest := &monitoringpb.ListTimeSeriesRequest{
		Name:   "projects/" + project,
		Filter: fmt.Sprintf(`metric.type = "compute.googleapis.com/quota/%s"`, quotaName),
		Interval: &monitoringpb.TimeInterval{
			StartTime: &timestamppb.Timestamp{
				Seconds: startTime.Unix(),
			},
			EndTime: &timestamppb.Timestamp{
				Seconds: endTime.Unix(),
			},
		},
	}
	return client.ListTimeSeries(ctx, monitoringRequest)
}

func loadCfg(cfgPath string) *gceproviderpb.Configs {
	in, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatalln("Error reading cfg file:", err)
	}
	configs := &gceproviderpb.Configs{}
	opts := prototext.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err := opts.Unmarshal(in, configs); err != nil {
		log.Fatalln("Failed to parse gce config:", err)
	}
	return configs
}

func zoneToRegion(zone string, possibleRegions []string) string {
	for _, region := range possibleRegions {
		if strings.HasPrefix(zone, region) {
			return region
		}
	}
	log.Fatalln("Couldn't find region for: ", zone)
	return zone
}

func getRegionQuotas(ctx context.Context, project string) (map[string]*regionQuotas, []string) {
	// Get regions and their quotas. Simply calling ListRegion will get
	// a variety of quota info for each region, which includes most of
	// what we care about.
	var regionNames []string
	quotasPerRegion := make(map[string]*regionQuotas)
	c, err := compute.NewRegionsRESTClient(ctx)
	if err != nil {
		log.Fatalln("Error init'ing gcloud client:", err)
	}
	defer func() {
		if err = c.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	req := &computepb.ListRegionsRequest{
		Project: project,
	}
	it := c.List(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln("Error listing regions:", err)
		}
		quotas := regionQuotas{localSSDPerFamilyQuota: make(map[string]*quotaVals)}
		for _, quota := range resp.Quotas {
			switch *quota.Metric {
			case "N2_CPUS":
				quotas.n2CpusQuota = quotaVals{max: int64(*quota.Limit)}
			case "CPUS":
				quotas.cpusQuota = quotaVals{max: int64(*quota.Limit)}
			case "IN_USE_ADDRESSES":
				quotas.ipsQuota = quotaVals{max: int64(*quota.Limit)}
			case "INSTANCES":
				quotas.instancesQuota = quotaVals{max: int64(*quota.Limit)}
			case "DISKS_TOTAL_GB":
				quotas.hddQuota = quotaVals{max: int64(*quota.Limit)}
			case "SSD_TOTAL_GB":
				quotas.remoteSSDQuota = quotaVals{max: int64(*quota.Limit)}
			}
		}
		quotasPerRegion[*resp.Name] = &quotas
		regionNames = append(regionNames, *resp.Name)
	}
	return quotasPerRegion, regionNames
}

func getLocalSSDQuotas(ctx context.Context, project string, quotasPerRegion map[string]*regionQuotas) {
	// Get local-SSD per region per vm-family quotas. Need to query a quota
	// metric for this. Note: we query the "limit" here and then calculate
	// the expected "usage" later on by reading gce-provider vms.cfg files.
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	timeSeriesIterator := queryTimeSeriesQuota(ctx, client, "local_ssd_total_storage_per_vm_family/limit", project)
	for {
		timeSeriesResp, err := timeSeriesIterator.Next()
		if err == iterator.Done {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		var machineFamily string
		var region string
		if value, ok := timeSeriesResp.Metric.Labels["vm_family"]; ok {
			machineFamily = strings.ToLower(value)
		}
		if value, ok := timeSeriesResp.Resource.Labels["location"]; ok {
			region = value
		}
		regionQuota, ok := quotasPerRegion[region]
		if !ok {
			continue
		}
		// Points[0] should correspond to the most recent value for the
		// quota, useful in case the quota limit gets bumped for the
		// project.
		regionQuota.localSSDPerFamilyQuota[machineFamily] = &quotaVals{max: timeSeriesResp.Points[0].Value.GetInt64Value()}
	}
}

func getNetworkQuotas(ctx context.Context, project string) map[string]*quotaVals {
	// Get the number of instances per network. Also need to query a quota
	// metric for this.
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = client.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	quotasPerNetwork := make(map[string]*quotaVals)
	timeSeriesIterator := queryTimeSeriesQuota(ctx, client, "instances_per_vpc_network/limit", project)
	for {
		timeSeriesResp, err := timeSeriesIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		var networkName string
		for labelKey, labelVal := range timeSeriesResp.Resource.Labels {
			// Need to map network ID to network name since GCE's
			// monitoring API uses the former while gce-provider
			// uses the latter.
			// FIXME: Get this mapping by querying the project?
			if labelKey == "network_id" {
				switch labelVal {
				case "655963314494161580":
					networkName = "c10"
				case "2688805488330601365":
					networkName = "c4"
				case "13012605346896030474":
					networkName = "default"
				case "2893051718470468954":
					networkName = "crbug1320004-test-network"
				default:
					log.Fatalln("Unknown network id: ", labelVal)
				}
				break
			}
		}
		quotasPerNetwork[networkName] = &quotaVals{max: timeSeriesResp.Points[0].Value.GetInt64Value()}
	}
	return quotasPerNetwork
}

func parseCfgFiles(project string, cfgPaths []string, regionNames []string, quotasPerRegion map[string]*regionQuotas, quotasPerNetwork map[string]*quotaVals) {
	// Read vms.cfg. The 'configs' slice is used to flatten all configs
	// across all possible cfg files.
	var configs []*gceproviderpb.Config
	for _, cfgPath := range cfgPaths {
		thisConfig := loadCfg(cfgPath)
		configs = append(configs, thisConfig.Vms...)
	}

	// Parse vms.cfg
	for _, config := range configs {
		if config.Attributes.Project != project {
			continue
		}
		region := zoneToRegion(config.Attributes.Zone, regionNames)

		// Get max num instances
		maxInstances := config.Amount.Max
		for _, scheduledChange := range config.Amount.Change {
			if scheduledChange.Max > maxInstances {
				maxInstances = scheduledChange.Max
			}
		}
		quotasPerRegion[region].instancesQuota.used += int64(maxInstances)

		// Get network
		if len(config.Attributes.NetworkInterface) != 1 || config.Attributes.NetworkInterface[0].Network == "" {
			log.Fatalln("Unknown network config on ", config.Prefix)
		}
		network := config.Attributes.NetworkInterface[0].Network
		network, _ = strings.CutPrefix(network, "global/networks/")
		quotasPerNetwork[network].used += int64(maxInstances)

		// TODO: Get IP address

		// TODO: Get core count

		// TODO: Get disk info
	}
}

func main() {
	ctx := context.Background()
	project, _, cfgPaths := parseFlags()

	// Query GCE for all relevant quotas for the project.
	quotasPerRegion, regionNames := getRegionQuotas(ctx, project)
	getLocalSSDQuotas(ctx, project, quotasPerRegion)
	quotasPerNetwork := getNetworkQuotas(ctx, project)

	// Parse vms.cfgs and determine max quota usage for all deployments for
	// the oroject.
	parseCfgFiles(project, cfgPaths, regionNames, quotasPerRegion, quotasPerNetwork)

	// TODO: Finish the rest.
}
