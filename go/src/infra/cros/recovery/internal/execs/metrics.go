// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package execs

import (
	"context"
	"time"

	"go.chromium.org/luci/common/errors"

	"infra/cros/recovery/internal/log"
	"infra/cros/recovery/logger/metrics"
)

// AddObservation adds observation to the metric assigned to the current exec.
func (i *ExecInfo) AddObservation(observation *metrics.Observation) {
	logger := i.NewLogger()
	if i.metric == nil {
		logger.Debugf("Metric is not specified for the action.")
	}
	logger.Infof("Add observation: %#v", observation)
	i.metric.Observations = append(i.metric.Observations, observation)
}

// NewMetric creates a new custom metric.
func (i *ExecInfo) NewMetric(kind string) *metrics.Action {
	// We do not check kind here as if it is empty then it will be rejected before saving.
	metric := i.runArgs.NewMetricsAction(kind)
	i.NewLogger().Debugf("Created new metrics for exec %q: %#v", i.name, metric)
	i.additionalMetrics = append(i.additionalMetrics, metric)
	return metric
}

// NewMetric creates a new metric.
func (a *RunArgs) NewMetricsAction(kind string) *metrics.Action {
	return &metrics.Action{
		// TODO(b/248635230): Set asset tag instead of hostname.
		Hostname:       a.DUT.Name,
		ActionKind:     kind,
		StartTime:      time.Now(),
		SwarmingTaskID: a.SwarmingTaskID,
		BuildbucketID:  a.BuildbucketID,
		Status:         metrics.ActionStatusUnspecified,
	}
}

// GetAdditionalMetrics returns additional metrics created by execs.
func (i *ExecInfo) GetAdditionalMetrics() []*metrics.Action {
	return i.additionalMetrics
}

// CloserFunc is a function that updates an action and is NOT safe to use in a defer block WITHOUT CHECKING FOR NIL.
// The following ways of using a CloserFunc are both correct.
//
// `ctx` and `err` are bound by the surrounding context.
//
//	action, closer := someFunction(...)
//	if closer != nil {
//	  defer closer(ctx, err)
//	}
//
//	action, closer := someFunction(...)
//	defer func() {
//	  if closer != nil {
//	    defer closer(ctx, err)
//	  }
//	}()
type CloserFunc = func(context.Context, error)

// NewMetric takes a reference to an action and populates it as a new action of kind `kind`.
// NewMetric mutates its argument action.
func (a *RunArgs) NewMetric(ctx context.Context, kind string, action *metrics.Action) (CloserFunc, error) {
	// Keep this function up to date with the call to args.Metrics.Create in recovery.go
	if a == nil {
		return nil, errors.Reason("new metrics: run args cannot be nil").Err()
	}
	if action == nil {
		return nil, errors.Reason("new metrics: action cannot be nil").Err()
	}
	dutName := ""
	if a.DUT != nil {
		dutName = a.DUT.Name
	}
	startTime := time.Now()
	*action = metrics.Action{
		ActionKind:     kind,
		StartTime:      startTime,
		SwarmingTaskID: a.SwarmingTaskID,
		BuildbucketID:  a.BuildbucketID,
		Hostname:       dutName,
	}
	c := createMetric(ctx, a.Metrics, action)
	return c, nil
}

// CreateMetric creates a metric with an actionKind, and a startTime.
// It returns an action and a closer function.
// CreateMetric mutates its argument action.
//
// Intended usage:
//
//	err is managed by the containing function.
//
//	Note that it is necessary to explicitly defer evaluation of err to the
//	end of the function.
//
//	closer := createMetric(ctx, ...)
//	if closer != nil {
//	  defer func() {
//	    closer(ctx, err)
//	  }()
//	}
func createMetric(ctx context.Context, m metrics.Metrics, action *metrics.Action) func(context.Context, error) {
	if m == nil {
		return nil
	}
	if err := m.Create(ctx, action); err != nil {
		log.Errorf(ctx, err.Error())
	}
	closer := func(ctx context.Context, e error) {
		if m == nil {
			log.Debugf(ctx, "Forgivable error while creating metric, nil metrics")
			return
		}
		if action == nil {
			log.Debugf(ctx, "Forgivable error while creating metric, action reference points to nil action")
			return
		}
		action.Status = metrics.ActionStatusUnspecified
		// TODO(gregorynisbet): Consider strategies for multiple fail reasons.
		if e != nil {
			log.Debugf(ctx, "Updating metric kind %q: marked as 'fail' with reason %q", action.ActionKind, e.Error())
			action.Status = metrics.ActionStatusFail
			action.FailReason = e.Error()
		} else {
			action.Status = metrics.ActionStatusSuccess
			log.Debugf(ctx, "Updating metric kind %q: close was successful", action.ActionKind)
		}
		if uErr := m.Update(ctx, action); uErr != nil {
			log.Errorf(ctx, "Updating metric kind %q: failed to update with error: %s", action.ActionKind, uErr.Error())
		}
	}
	return closer
}
