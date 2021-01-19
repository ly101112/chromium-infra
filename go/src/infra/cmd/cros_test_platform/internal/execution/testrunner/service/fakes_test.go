// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package service_test

import (
	"context"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"go.chromium.org/luci/common/data/stringset"

	trservice "infra/cmd/cros_test_platform/internal/execution/testrunner/service"
	"infra/libs/skylab/inventory"
	"infra/libs/skylab/request"
)

func TestBotsAwareFakeClient(t *testing.T) {
	cases := []struct {
		Tag          string
		Client       trservice.BotsAwareFakeClient
		Args         request.Args
		WantValid    bool
		WantRejected []string
	}{
		{
			Tag:    "no bots with free-form dimensions in build",
			Client: trservice.NewBotsAwareFakeClient(),
			Args: request.Args{
				Dimensions: []string{"free-form:value"},
			},
			WantValid:    false,
			WantRejected: []string{"free-form"},
		},
		{
			Tag:    "no bots with schedulable labels in build",
			Client: trservice.NewBotsAwareFakeClient(),
			Args: request.Args{
				SchedulableLabels: &inventory.SchedulableLabels{
					Board: stringPtr("foo"),
				},
			},
			WantValid:    false,
			WantRejected: []string{"label-board"},
		},
		{
			Tag:    "mismatched bot",
			Client: trservice.NewBotsAwareFakeClient(stringset.NewFromSlice("free-form:bot-value")),
			Args: request.Args{
				Dimensions: []string{"free-form:build-value"},
			},
			WantValid:    false,
			WantRejected: []string{"free-form"},
		},
		{
			Tag:    "matched bot",
			Client: trservice.NewBotsAwareFakeClient(stringset.NewFromSlice("free-form:bot-value")),
			Args: request.Args{
				Dimensions: []string{"free-form:bot-value"},
			},
			WantValid:    true,
			WantRejected: nil,
		},
	}
	for _, c := range cases {
		t.Run(c.Tag, func(t *testing.T) {
			b, r, err := c.Client.ValidateArgs(context.Background(), &c.Args)
			if err != nil {
				t.Fatalf("ValidateArgs returned error: %s", err)
			}
			if b != c.WantValid {
				t.Errorf("ValidateArgs returned %t, want %t", b, c.WantValid)
			}
			for _, k := range c.WantRejected {
				if _, ok := r[k]; !ok {
					t.Errorf("Rejected arguments missing key %s", k)
				}
			}
		})
	}
}

func TestClientCallCountingWrapper(t *testing.T) {
	c := trservice.ClientCallCountingWrapper{
		Client: trservice.StubClient{},
	}
	if diff := pretty.Compare(trservice.ClientMethodCallCounts{}, c.MethodCallCounts()); diff != "" {
		t.Fatalf("precondition counts differ, -want, +got: %s", diff)
	}

	c.ValidateArgs(context.Background(), nil)
	c.LaunchTask(context.Background(), nil)
	c.FetchResults(context.Background(), "")
	c.SwarmingTaskID("")
	c.URL("")
	want := trservice.ClientMethodCallCounts{
		ValidateArgs:   1,
		LaunchTask:     1,
		FetchResults:   1,
		SwarmingTaskID: 1,
		URL:            1,
	}
	if diff := pretty.Compare(want, c.MethodCallCounts()); diff != "" {
		t.Fatalf("post-condition counts differ, -want, +got: %s", diff)
	}
}

func stringPtr(val string) *string {
	return &val
}
