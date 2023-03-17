// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package build

import (
	"container/heap"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"infra/build/siso/ui"
)

type progress struct {
	mu sync.Mutex
	ts time.Time

	actives activeSteps
	done    chan struct{}
}

type stepInfo struct {
	step *Step
	desc string
}

type activeSteps []*stepInfo

func (as activeSteps) Len() int           { return len(as) }
func (as activeSteps) Less(i, j int) bool { return as[i].step.startTime.Before(as[j].step.startTime) }
func (as activeSteps) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }
func (as *activeSteps) Push(x interface{}) {
	(*as) = append(*as, x.(*stepInfo))
}
func (as *activeSteps) Pop() interface{} {
	old := *as
	n := len(old)
	s := old[n-1]
	old[n-1] = nil
	*as = old[:n-1]
	return s
}

func (p *progress) start(ctx context.Context, b *Builder) {
	p.done = make(chan struct{})
	if ui.IsTerminal {
		go p.update(ctx, b)
	}
}

func (p *progress) update(ctx context.Context, b *Builder) {
	lastUpdate := time.Now()
	lastStepUpdate := time.Now()
	for {
		select {
		case <-p.done:
			return
		case <-ctx.Done():
			return
		case <-time.After(100 * time.Millisecond):
			p.mu.Lock()
			var si *stepInfo
			for len(p.actives) > 0 {
				s := p.actives[0]
				if s.step.Done() {
					// already finished?
					heap.Pop(&p.actives)
					continue
				}
				si = s
				break
			}
			if len(p.actives) > 0 {
				d := time.Since(lastUpdate)
				wd := d / time.Duration(len(p.actives))
				lastUpdate = time.Now()
				for _, s := range p.actives {
					s.step.addWeightedDuration(wd)
				}
			}
			p.mu.Unlock()
			if si == nil || si.step == nil {
				continue
			}
			if time.Since(lastStepUpdate) < 1*time.Second {
				continue
			}
			lastStepUpdate = time.Now()
			dur := time.Since(si.step.startTime).Round(1 * time.Second)
			p.step(ctx, b, si.step, fmt.Sprintf("%s[%s]: %s", dur.String(), si.step.Phase(), si.desc))
		}
	}
}

func (p *progress) stop(ctx context.Context) {
	close(p.done)
}

func (p *progress) report(format string, args ...interface{}) {
	p.mu.Lock()
	t := p.ts
	p.mu.Unlock()
	if ui.IsTerminal && time.Since(t) < 500*time.Millisecond {
		return
	}
	msg := fmt.Sprintf(format, args...)
	ui.PrintLines(msg)
	p.mu.Lock()
	p.ts = time.Now()
	p.mu.Unlock()
}

const (
	progressPrefixStart  = "S "
	progressPrefixFinish = "F "
)

func (p *progress) step(ctx context.Context, b *Builder, step *Step, s string) {
	p.mu.Lock()
	t := p.ts
	if step != nil && ui.IsTerminal {
		if strings.HasPrefix(s, progressPrefixStart) {
			heap.Push(&p.actives, &stepInfo{
				step: step,
				desc: step.cmd.Desc,
			})
		}
	}
	p.mu.Unlock()
	if ui.IsTerminal && (time.Since(t) < 100*time.Millisecond || strings.HasPrefix(s, progressPrefixFinish)) {
		return
	}
	pstat := b.plan.stats()
	stat := b.stats.stats()
	msg := s
	if step != nil {
		msg = fmt.Sprintf("%.2f%% %s %s",
			float64(stat.Done)*100/float64(stat.Total),
			time.Since(b.start).Round(10*time.Millisecond),
			s)
	}
	ui.PrintLines(
		fmt.Sprintf("[%d+%d+(%d+%d)[%d+(%d+%d),(%d+%d)](%d,%d,%d,%d){%d,%d,%d}]",
			pstat.npendings,
			pstat.nready,
			b.stepSema.NumWaits(),
			b.stepSema.NumServs(),
			stat.Preproc,
			b.localSema.NumWaits(),
			b.localSema.NumServs(),
			b.remoteSema.NumWaits(),
			b.remoteSema.NumServs(),
			stat.Skipped,
			stat.CacheHit,
			stat.Local,
			stat.Remote,
			stat.FastDepsSuccess,
			stat.FastDepsFailed,
			stat.LocalFallback),
		msg,
	)
	p.mu.Lock()
	p.ts = time.Now()
	p.mu.Unlock()

}