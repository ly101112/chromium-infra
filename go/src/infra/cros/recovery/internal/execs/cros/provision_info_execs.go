// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package cros

import (
	"context"
	"fmt"

	"go.chromium.org/luci/common/errors"

	"infra/cros/recovery/internal/components/cros"
	"infra/cros/recovery/internal/execs"
	"infra/cros/recovery/internal/log"
	"infra/cros/recovery/tlw"
)

const (
	// gsCrOSImageBucket is the base URL for the Google Storage bucket for
	// ChromeOS image archives.
	gsCrOSImageBucket = "gs://chromeos-image-archive"
)

// updateProvisionedInfoExec updated provision info.
// Data which included:
// 1) OS version from the DUT.
// 2) Job repo URL to download the packages matched to OS version.
func updateProvisionedInfoExec(ctx context.Context, info *execs.ExecInfo) error {
	if info.GetDut().ProvisionedInfo == nil {
		info.GetDut().ProvisionedInfo = &tlw.ProvisionedInfo{}
	}
	run := info.NewRunner(info.GetDut().Name)
	osVersion, err := cros.ReleaseBuildPath(ctx, run, info.NewLogger())
	if err != nil {
		return errors.Annotate(err, "update provision info").Err()
	}
	log.Debugf(ctx, "ChromeOS version on the dut: %s.", osVersion)
	info.GetDut().ProvisionedInfo.CrosVersion = osVersion

	argsMap := info.GetActionArgs(ctx)
	if argsMap.AsBool(ctx, "update_job_repo_url", false) {
		gsPath := fmt.Sprintf("%s/%s", gsCrOSImageBucket, osVersion)
		jobRepoURL, err := info.GetAccess().GetCacheUrl(ctx, info.GetDut().Name, gsPath)
		if err != nil {
			return errors.Annotate(err, "update provision info").Err()
		}
		log.Debugf(ctx, "New job repo URL: %s.", jobRepoURL)
		info.GetDut().ProvisionedInfo.JobRepoUrl = jobRepoURL
	}
	return nil
}

func init() {
	execs.Register("cros_update_provision_info", updateProvisionedInfoExec)
}