// Copyright 2024 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package costserver_test

import (
	"context"
	"testing"

	fleetcostpb "infra/cros/fleetcost/api"
	testsupport "infra/cros/fleetcost/internal/costserver/testsupport"
)

func TestCreateCostIndicator(t *testing.T) {
	t.Parallel()
	tf := testsupport.NewFixture(context.Background(), t)

	_, err := tf.Frontend.CreateCostIndicator(tf.Ctx, &fleetcostpb.CreateCostIndicatorRequest{})
	if err == nil {
		t.Error("err is unexpectedly nil")
	}
}