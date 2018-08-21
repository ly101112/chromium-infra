// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scheduler

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"infra/qscheduler/qslib/mutaters"
	"infra/qscheduler/qslib/types"
	"infra/qscheduler/qslib/types/account"
	"infra/qscheduler/qslib/types/task"
	"infra/qscheduler/qslib/types/vector"

	"github.com/golang/protobuf/ptypes"
)

// epoch is an arbitrary time for testing purposes, corresponds to
// 01/01/2018 @ 1:00 am UTC
var epoch = time.Unix(1514768400, 0)

// assertMutations is a testing helper that does a itemwise comparison of two
// slices of mutaters.Mutater, and fails if they are unequal.
func assertMutations(t *testing.T, expects []mutaters.Mutater, actual []mutaters.Mutater, caseNum int) {
	if len(actual) != len(expects) {
		t.Errorf("Case %d got %d muts \n%+v\n, want %d muts \n%+v", caseNum, len(actual), actual, len(expects), expects)
		return
	}

	for i, mut := range actual {
		expect := expects[0]
		expects = expects[1:]
		if !reflect.DeepEqual(mut, expect) {
			t.Errorf("Case %d at element %d got: %+v, want: %+v", caseNum, i, mut, expect)
		}
	}
}

func TestMatchWithIdleWorkers(t *testing.T) {
	t.Parallel()
	state := types.State{
		Workers: map[string]*types.Worker{
			"w0": types.NewWorker(),
			"w1": &types.Worker{Labels: []string{"label1"}},
		},
		Requests: map[string]*task.Request{
			"t1": &task.Request{AccountId: "a1", Labels: []string{"label1"}},
			"t2": &task.Request{AccountId: "a1", Labels: []string{"label2"}},
		},
		Balances: map[string]*vector.Vector{
			"a1": vector.New(2, 0, 0),
		},
	}

	config := types.Config{
		AccountConfigs: map[string]*account.Config{
			"a1": account.NewConfig(),
		},
	}

	expects := []mutaters.Mutater{
		&mutaters.AssignIdleWorker{Priority: 0, RequestId: "t1", WorkerId: "w1"},
		&mutaters.AssignIdleWorker{Priority: 0, RequestId: "t2", WorkerId: "w0"},
	}

	muts := QuotaSchedule(&state, &config)
	assertMutations(t, expects, muts, 0)
}

func TestReprioritize(t *testing.T) {
	t.Parallel()
	// Prepare a situation in which one P0 job (out of 2 running) will be
	// demoted, and a separate P2 job will be promoted to P1.
	config := types.Config{
		AccountConfigs: map[string]*account.Config{
			"a1": &account.Config{ChargeRate: vector.New(1.5, 1.5)},
		},
	}
	state := types.State{
		Balances: map[string]*vector.Vector{
			"a1": vector.New(2*account.DemoteThreshold, 2*account.PromoteThreshold, 0),
		},
		Workers: map[string]*types.Worker{
			"w1": &types.Worker{
				RunningTask: &task.Run{
					Cost:     vector.New(1),
					Priority: 0,
					Request:  &task.Request{AccountId: "a1"},
				},
			},
			"w2": &types.Worker{
				RunningTask: &task.Run{
					Priority: 0,
					Request:  &task.Request{AccountId: "a1"},
					Cost:     vector.New(),
				},
			},
			"w3": &types.Worker{
				RunningTask: &task.Run{
					Cost:     vector.New(1),
					Priority: 2,
					Request:  &task.Request{AccountId: "a1"},
				},
			},
			"w4": &types.Worker{
				RunningTask: &task.Run{
					Priority: 2,
					Request:  &task.Request{AccountId: "a1"},
					Cost:     vector.New(),
				},
			},
		},
	}

	expects := []mutaters.Mutater{
		&mutaters.ChangePriority{NewPriority: 1, WorkerId: "w2"},
		&mutaters.ChangePriority{NewPriority: 1, WorkerId: "w3"},
	}

	muts := QuotaSchedule(&state, &config)
	assertMutations(t, expects, muts, 0)
}

func TestPreempt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		State  *types.State
		Config *types.Config
		Expect []mutaters.Mutater
	}{
		// Case 0
		//
		// Basic preemption of a job by a higher priority job.
		{
			&types.State{
				Balances: map[string]*vector.Vector{
					"a1": vector.New(),
					"a2": vector.New(1),
				},
				Requests: map[string]*task.Request{
					"t1": &task.Request{AccountId: "a2"},
				},
				Workers: map[string]*types.Worker{
					"w1": &types.Worker{
						RunningTask: &task.Run{
							Cost:      vector.New(.5, .5, .5),
							Priority:  1,
							Request:   &task.Request{AccountId: "a1"},
							RequestId: "t2",
						},
					},
				},
			},
			&types.Config{
				AccountConfigs: map[string]*account.Config{
					"a1": account.NewConfig(),
					"a2": account.NewConfig(),
				},
			},
			[]mutaters.Mutater{&mutaters.PreemptTask{Priority: 0, WorkerId: "w1", RequestId: "t1"}},
		},
		// Case 1
		//
		// Preemption will be skipped if:
		// - the preempting account has insufficient funds.
		// - the preempting account already has lower priority jobs.
		{
			&types.State{
				// Both accounts a1 and a2 have P0 quota.
				Balances: map[string]*vector.Vector{
					// a1 has insufficient balance to preempt jobs.
					"a1": vector.New(0.1 * account.PromoteThreshold),
					// a2 would have sufficient balance to preempt jobs, but has
					// insufficient balance to promote its already running job, and
					// thus is banned from preempting jobs.
					"a2": vector.New(0.9 * account.PromoteThreshold),
				},
				Requests: map[string]*task.Request{
					"t1": &task.Request{AccountId: "a1"},
					"t2": &task.Request{AccountId: "a2"},
				},
				Workers: map[string]*types.Worker{
					// A job is running, but it is too costly for a1 to preempt.
					"w1": &types.Worker{
						RunningTask: &task.Run{
							Cost:      vector.New(0.5*account.PromoteThreshold, 0, 0),
							Priority:  1,
							Request:   &task.Request{},
							RequestId: "other_req",
						},
					},
					// A job is running for a2 at a lower priority, so a2 is banned
					// from preempting jobs.
					"w2": &types.Worker{
						RunningTask: &task.Run{
							Cost:      vector.New(0.5 * account.PromoteThreshold),
							Priority:  1,
							Request:   &task.Request{AccountId: "a2"},
							RequestId: "t3",
						},
					},
				},
			},
			&types.Config{
				AccountConfigs: map[string]*account.Config{
					"a1": &account.Config{ChargeRate: vector.New(1)},
					"a2": &account.Config{ChargeRate: vector.New(1)},
				},
			},
			// No preemptions or other mutations should result.
			[]mutaters.Mutater{},
		},
	}

	for i, test := range cases {
		actual := QuotaSchedule(test.State, test.Config)
		assertMutations(t, test.Expect, actual, i)
	}
}

func TestUpdateErrors(t *testing.T) {
	cases := []struct {
		State  *types.State
		Config *types.Config
		T      time.Time
		Expect error
	}{
		{
			types.NewState(),
			types.NewConfig(),
			time.Unix(0, 0),
			errors.New("timestamp: nil Timestamp"),
		},
		{
			stateAtTime(time.Unix(100, 0).UTC()),
			types.NewConfig(),
			time.Unix(0, 0).UTC(),
			&UpdateOrderError{Next: time.Unix(0, 0).UTC(), Previous: time.Unix(100, 0).UTC()},
		},
		{
			stateAtTime(time.Unix(0, 0)),
			types.NewConfig(),
			time.Unix(1, 0),
			nil,
		},
	}

	for i, test := range cases {
		e := UpdateAccounts(test.State, test.Config, test.T)
		if !reflect.DeepEqual(e, test.Expect) {
			t.Errorf("In case %d, got error: %+v, want error: %+v", i, e, test.Expect)
		}
	}
}

func TestUpdateBalance(t *testing.T) {
	t0, _ := ptypes.TimestampProto(epoch)
	t1, _ := ptypes.TimestampProto(epoch.Add(1 * time.Second))
	t2, _ := ptypes.TimestampProto(epoch.Add(2 * time.Second))

	cases := []struct {
		State  *types.State
		Config *types.Config
		T      time.Time
		Expect *types.State
	}{
		// Case 0:
		// Balances with no account config should be removed ("a1"). New balances
		// should be created if necessary and incremented appropriately ("a2").
		{
			&types.State{
				Balances:          map[string]*vector.Vector{"a1": vector.New()},
				LastAccountUpdate: t0,
			},
			&types.Config{
				AccountConfigs: map[string]*account.Config{
					"a2": &account.Config{ChargeRate: vector.New(1), MaxChargeSeconds: 2},
				},
			},
			epoch.Add(1 * time.Second),
			&types.State{
				Balances:          map[string]*vector.Vector{"a2": vector.New(1)},
				LastAccountUpdate: t1,
			},
		},
		// Case 1:
		// Running jobs should count against the account. Cost of a running job
		// should be initialized if necessary, and incremented.
		//
		// Charges should be proportional to time advanced (2 seconds in this case).
		{
			&types.State{
				Balances: map[string]*vector.Vector{"a1": vector.New()},
				Workers: map[string]*types.Worker{
					// Worker running a task.
					"w1": &types.Worker{
						RunningTask: &task.Run{
							Cost:     vector.New(1),
							Priority: 1,
							Request:  &task.Request{AccountId: "a1"},
						},
					},
					// Worker running a task with uninitialized Cost.
					"w2": &types.Worker{
						RunningTask: &task.Run{
							Priority: 2,
							Request:  &task.Request{AccountId: "a1"},
						},
					},
					// Worker running a task with invalid account.
					"w3": &types.Worker{
						RunningTask: &task.Run{
							Priority: account.FreeBucket,
							Request:  &task.Request{AccountId: "a2"},
						},
					},
				},
				LastAccountUpdate: t0,
			},
			&types.Config{
				AccountConfigs: map[string]*account.Config{
					"a1": &account.Config{ChargeRate: vector.New(1), MaxChargeSeconds: 1},
				},
			},
			epoch.Add(2 * time.Second),
			&types.State{
				Balances:          map[string]*vector.Vector{"a1": vector.New(1, -2, -2)},
				LastAccountUpdate: t2,
				Workers: map[string]*types.Worker{
					"w1": &types.Worker{
						RunningTask: &task.Run{
							Cost:     vector.New(1, 2),
							Priority: 1,
							Request:  &task.Request{AccountId: "a1"},
						},
					},
					"w2": &types.Worker{
						RunningTask: &task.Run{
							Cost:     vector.New(0, 0, 2),
							Priority: 2,
							Request:  &task.Request{AccountId: "a1"},
						},
					},
					"w3": &types.Worker{
						RunningTask: &task.Run{
							Cost:     vector.New(),
							Priority: account.FreeBucket,
							Request:  &task.Request{AccountId: "a2"},
						},
					},
				},
			},
		},
	}

	for i, test := range cases {
		actual := test.State
		UpdateAccounts(actual, test.Config, test.T)
		if !reflect.DeepEqual(actual, test.Expect) {
			t.Errorf("In case %d got state\n%+v\nwant\n%+v", i, actual, test.Expect)
		}
	}
}

func stateAtTime(t time.Time) *types.State {
	s := types.NewState()
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		panic(err)
	}
	s.LastAccountUpdate = ts
	return s
}
