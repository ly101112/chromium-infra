// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package wifirouter

import (
	"go.chromium.org/luci/common/errors"
	"infra/cros/recovery/internal/execs"
	"infra/cros/recovery/internal/execs/wifirouter/controller"
	"infra/cros/recovery/tlw"
)

// activeHost finds active host related to the executed plan.
func activeHost(info *execs.ExecInfo) (*tlw.WifiRouterHost, error) {
	resource := info.GetActiveResource()
	chromeos := info.GetChromeos()
	if chromeos == nil {
		return nil, errors.Reason("chromeos is not present").Err()
	}
	for _, router := range chromeos.GetWifiRouters() {
		if router.GetName() == resource {
			return router, nil
		}
	}
	return nil, errors.Reason("router: router not found").Err()
}

func activeHostRouterController(info *execs.ExecInfo) (controller.RouterController, error) {
	wifiRouterHost, err := activeHost(info)
	if err != nil {
		return nil, err
	}
	return controller.NewRouterDeviceController(info.GetAccess(), info.GetActiveResource(), wifiRouterHost)
}
