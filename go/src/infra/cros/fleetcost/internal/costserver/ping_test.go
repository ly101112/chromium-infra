// Copyright 2024 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package costserver

import (
	"context"
	"testing"

	fleetcostpb "infra/cros/fleetcost/api"
)

// TestPing tests the ping API, which does nothing
func TestPing(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	tf := newFixture(ctx, t)

	_, err := tf.frontend.Ping(tf.ctx, &fleetcostpb.PingRequest{})
	if err != nil {
		t.Error(err)
	}
}
