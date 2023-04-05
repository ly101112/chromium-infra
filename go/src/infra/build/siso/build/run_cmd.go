// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package build

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/golang/glog"

	"infra/build/siso/o11y/clog"
	"infra/build/siso/o11y/trace"
	"infra/build/siso/reapi"
)

// runCmdWithCache checks and returns the ActionResult cache,
// or runs the command if the cache doesn't exist.
func (b *Builder) runCmdWithCache(ctx context.Context, step *Step, allowLocalFallback bool) error {
	dedupInputs(ctx, step.Cmd)
	if b.cache != nil && step.Cmd.Pure && b.reCacheEnableRead {
		err := b.cacheSema.Do(ctx, func(ctx context.Context) error {
			start := time.Now()
			err := b.cache.GetActionResult(ctx, step.Cmd)
			if err != nil {
				return err
			}
			err = b.outputs(ctx, step)
			if err != nil {
				return err
			}
			b.stats.cacheHit(ctx)
			b.progressStepCacheHit(ctx, step)
			step.Metrics.Cached = true
			step.Metrics.RunTime = IntervalMetric(time.Since(start))
			step.Metrics.Done(ctx, step)
			// need to update deps for cache hit for deps=gcc, msvc.
			// even if cache hit, deps should be updated with gcc depsfile,
			// or with msvc showIncludes outputs.
			err = b.updateDeps(ctx, step)
			if err != nil {
				clog.Warningf(ctx, "failed to update deps %s: %v", step, err)
			}
			return nil
		})
		if err == nil {
			return nil
		}
		clog.Infof(ctx, "cmd cache miss: %v", err)
	}
	err := b.runCmd(ctx, step, allowLocalFallback)
	if err != nil {
		if !allowLocalFallback {
			clog.Warningf(ctx, "step fail->!localFallback")
			return fmt.Errorf("run for %s [no allowLocalFallback]: %w", step, err)
		}
		return fmt.Errorf("run for %s: %w", step, err)
	}
	// TODO(b/266518906): enable reCacheEnableWrite option for read-only client.
	return b.outputs(ctx, step)
}

// runCmd runs the command remotely if step is remote executable and remoteexec is configured,
// or locally if step is not remote executable, remoteexec is not configured or failed and allowLocalFallback is true.
func (b *Builder) runCmd(ctx context.Context, step *Step, allowLocalFallback bool) error {
	ctx, span := trace.NewSpan(ctx, "run")
	defer span.Close(nil)
	if log.V(1) {
		clog.Infof(ctx, "run %s [allow-localfallback=%t]", step.Cmd.Desc, allowLocalFallback)
	}
	if step.Cmd.Pure && len(step.Cmd.Platform) > 0 && step.Cmd.Platform["container-image"] != "" && b.remoteExec != nil {
		err := b.runRemote(ctx, step)
		if err == nil {
			// need to check remote outptus matches cmd.Outputs?
			return nil
		}
		if errors.Is(err, context.Canceled) {
			return err
		}
		if errors.Is(err, reapi.ErrBadPlatformContainerImage) {
			return err
		}
		if !allowLocalFallback {
			return fmt.Errorf("no allow-localfallback: %w", err)
		}
		if errors.Is(err, errNotRelocatable) {
			clog.Errorf(ctx, "not relocatable: %v", err)
			return err
		}
		b.stats.localFallback(ctx)
		if experiments.Enabled("no-fallback", "remote-exec %s failed. no-fallback", step) {
			return fmt.Errorf("remote-exec %s failed. no-fallback: %w", step, err)
		}
		step.Metrics.Fallback = true
	}
	if !allowLocalFallback {
		return errors.New("no allow-localfallback")
	}
	return b.runLocal(ctx, step)
}