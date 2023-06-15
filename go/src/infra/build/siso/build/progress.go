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
	"sync/atomic"
	"time"

	"infra/build/siso/ui"
)

type progress struct {
	mu sync.Mutex
	ts time.Time

	actives       activeSteps
	done          chan struct{}
	updateStopped chan struct{}
	count         atomic.Int64
}

type stepInfo struct {
	step *Step
	desc string
}

type activeSteps []*stepInfo

func (as activeSteps) Len() int           { return len(as) }
func (as activeSteps) Less(i, j int) bool { return as[i].step.startTime.Before(as[j].step.startTime) }
func (as activeSteps) Swap(i, j int)      { as[i], as[j] = as[j], as[i] }
func (as *activeSteps) Push(x any) {
	(*as) = append(*as, x.(*stepInfo))
}
func (as *activeSteps) Pop() any {
	old := *as
	n := len(old)
	s := old[n-1]
	old[n-1] = nil
	*as = old[:n-1]
	return s
}

func (p *progress) start(ctx context.Context, b *Builder) {
	p.done = make(chan struct{})
	p.updateStopped = make(chan struct{})
	go p.update(ctx, b)
}

func (p *progress) update(ctx context.Context, b *Builder) {
	lastUpdate := time.Now()
	lastStepUpdate := time.Now()
	defer close(p.updateStopped)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-p.done:
			return
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.count.Add(1)
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
			if !ui.IsTerminal() {
				continue
			}
			if si == nil || si.step == nil {
				continue
			}
			const d = 100 * time.Millisecond
			if time.Since(lastStepUpdate) < d {
				continue
			}
			lastStepUpdate = time.Now()
			dur := time.Since(si.step.startTime).Round(d)
			p.step(ctx, b, si.step, fmt.Sprintf("%s[%s]: %s", formatDuration(dur), si.step.phase(), si.desc))
		}
	}
}

func (p *progress) stop(ctx context.Context) {
	close(p.done)
	<-p.updateStopped
}

func (p *progress) report(format string, args ...any) {
	p.mu.Lock()
	t := p.ts
	p.mu.Unlock()
	if ui.IsTerminal() && time.Since(t) < 500*time.Millisecond {
		return
	}
	msg := fmt.Sprintf(format, args...)
	ui.Default.PrintLines(msg)
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
	if step != nil {
		if strings.HasPrefix(s, progressPrefixStart) {
			heap.Push(&p.actives, &stepInfo{
				step: step,
				desc: step.cmd.Desc,
			})
		}
	}
	p.mu.Unlock()
	if ui.IsTerminal() && (time.Since(t) < 30*time.Millisecond || strings.HasPrefix(s, progressPrefixFinish)) {
		return
	}
	pstat := b.plan.stats()
	stat := b.stats.stats()
	var lines []string
	if ui.IsTerminal() {
		lines = append(lines, fmt.Sprintf("[%d+%d+(%d+%d)[%d+(%d+%d),(%d+%d)](%d,%d,%d,%d){%d,%d,%d}]",
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
			stat.LocalFallback))
	}
	msg := s
	if step != nil {
		msg = fmt.Sprintf("%.2f%% %s %s",
			float64(stat.Done)*100/float64(stat.Total),
			formatDuration(time.Since(b.start)),
			s)
	}
	lines = append(lines, msg)
	ui.Default.PrintLines(lines...)
	p.mu.Lock()
	p.ts = time.Now()
	p.mu.Unlock()
}

func formatDuration(d time.Duration) string {
	d = d.Round(10 * time.Millisecond)
	var sb strings.Builder
	sb.Grow(32)

	mins := d.Truncate(1 * time.Minute)
	d = d - mins
	if mins > 0 {
		fmt.Fprintf(&sb, "%s", strings.TrimSuffix(mins.String(), "0s"))
		if d < 10*time.Second {
			fmt.Fprint(&sb, "0")
		}
	}
	fmt.Fprintf(&sb, "%02.02fs", d.Seconds())
	return sb.String()
}
