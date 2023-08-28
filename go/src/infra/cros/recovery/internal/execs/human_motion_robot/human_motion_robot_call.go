// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package human_motion_robot

import (
	"context"

	"go.chromium.org/chromiumos/config/go/api/test/xmlrpc"
	"go.chromium.org/luci/common/errors"

	"infra/cros/recovery/internal/log"
	"infra/cros/recovery/tlw"
)

// Call calls XMLRPC on touchhost.
func Call(ctx context.Context, in tlw.Access, host *tlw.HumanMotionRobot, method string) (*xmlrpc.Value, error) {
	if method == "" {
		return nil, errors.Reason("HMR TouchHost call: method name is empty").Err()
	}
	res := in.CallTouchHostd(ctx, &tlw.CallTouchHostdRequest{
		Resource: host.Name,
		Method:   method,
	})
	log.Debugf(ctx, "HMR TouchHost call %q: received %#v", method, res.GetValue())
	if res.GetFault() {
		return nil, errors.Reason("HMR TouchHost call %q: received fail flag. The message: %#v", method, res.GetValue()).Err()
	}
	return res.GetValue(), nil
}
