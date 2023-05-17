// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package build

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

	log "github.com/golang/glog"

	"infra/build/siso/hashfs"
	"infra/build/siso/o11y/clog"
	"infra/build/siso/ui"
)

var (
	// ErrNoTarget is an error when target is not found in the Graph.
	ErrNoTarget = errors.New("target not found")

	// ErrTargetIsSource is an error when target is the source in the Graph.
	ErrTargetIsSource = errors.New("target is source")

	// ErrDuplicateStep is an error when step for the target is already created.
	ErrDuplicateStep = errors.New("duplicate step")

	// ErrMissingDeps is an error when step failed to get deps.
	ErrMissingDeps = errors.New("missing deps in depfile/depslog")
)

// Graph provides a build graph, i.e. step definitions.
type Graph interface {
	// Targets returns target paths for given args.
	// Targets are exec-root relative.
	Targets(context.Context, ...string) ([]string, error)

	// StepDef creates new StepDef for the target and its inputs.
	// Target and inputs should be exec-root relative.
	// if err is ErrTargetIsSource, target is source and no step to
	// generate the target.
	// if err is ErrDuplicateStep, a step that geneartes the target
	// is already processed.
	StepDef(context.Context, string, StepDef) (StepDef, []string, error)

	// InputDeps returns input dependencies.
	// input dependencies is a map from input path or label to
	// other files or labels needed for the key.
	// path is exec root relative and label contains ':'.
	// it's "input_deps" in starlark config.
	InputDeps(context.Context) map[string][]string
}

// plan maintains which step to execute next.
type plan struct {
	// marked source target
	m map[string]bool

	mu        sync.Mutex
	q         chan *Step
	closed    bool
	ready     []*Step
	waits     map[string][]*Step
	npendings int
}

// schedulerOption is scheduler option.
type schedulerOption struct {
	Path        *Path
	HashFS      *hashfs.HashFS
	EnableTrace bool
}

// scheduler creates a plan.
type scheduler struct {
	path   *Path
	hashFS *hashfs.HashFS

	plan  *plan
	stats *stats

	lastProgress time.Time
	visited      int

	enableTrace bool
}

// schedule schedules build plans for args from graph into sched.
func schedule(ctx context.Context, sched *scheduler, graph Graph, args ...string) error {
	started := time.Now()
	targets, err := graph.Targets(ctx, args...)
	clog.Infof(ctx, "schedule targets: %q: %v", targets, err)
	if err != nil {
		return err
	}
	for _, t := range targets {
		err := scheduleTarget(ctx, sched, graph, t, nil)
		if err != nil {
			return fmt.Errorf("failed in schedule %s: %w", t, err)
		}
	}
	sched.finish(ctx, time.Since(started))
	return nil
}

// scheduleTarget schedules a build plan for target, which is required to next StepDef, from graph into sched.
func scheduleTarget(ctx context.Context, sched *scheduler, graph Graph, target string, next StepDef) error {
	if sched.marked(target) {
		if log.V(1) {
			clog.Infof(ctx, "sched target already marked: %s", target)
		}
		return nil
	}
	if log.V(1) {
		clog.Infof(ctx, "schedule target %s", target)
	}
	newStep, inputs, err := graph.StepDef(ctx, target, next)
	switch {
	case err == nil:
		// need to schedule.
	case errors.Is(err, ErrNoTarget):
		if log.V(1) {
			clog.Infof(ctx, "sched target not found? %s", target)
		}
		return err
	case errors.Is(err, ErrTargetIsSource):
		if log.V(1) {
			clog.Infof(ctx, "sched target is source? %s", target)
		}
		return sched.mark(ctx, target)
	case errors.Is(err, ErrDuplicateStep):
		// this step is already processed.
		if log.V(1) {
			clog.Infof(ctx, "sched duplicate step for %s", target)
		}
		return nil
	default:
		if log.V(1) {
			clog.Warningf(ctx, "sched error for %s: %v", target, err)
		}
		return err
	}
	clog.Infof(ctx, "schedule %s", newStep)
	sched.visited++
	next = newStep
	select {
	case <-ctx.Done():
		return fmt.Errorf("interrupted in schedule: %w", ctx.Err())
	default:
	}
	// to suppress duplicates
	m := make(map[string]bool)

	// we might not need to use depfile's dependencies to construct
	// build graph.
	// - if depfile's dependency is source file, the file already exists
	//   so no need to wait for it. doesn't change build graph.
	// - if depfile's dependency is generated file
	//   - and if it is included in step's inputs, or indirect inputs,
	//     then, it just adds redundant edge to build graph. Without
	//     the edge, step's order won't be changed, so no need to add
	//     such edge.
	//   - otherwise, it will change the build graph.
	//     it means original build graph without depfile contains
	//     missing dependencies. It would be better to fix gn/ninja's
	//     build graph, rather than mitigating here in the siso.
	var waits []string
	for _, in := range inputs {
		if m[in] {
			continue
		}
		m[in] = true
		err := scheduleTarget(ctx, sched, graph, in, next)
		if err != nil {
			return fmt.Errorf("schedule %s: %w", in, err)
		}
		if !sched.marked(in) {
			// If in is not marked (i.e. source), some step
			// will generate it, so need to wait for it
			// before running this step.
			waits = append(waits, in)
		}
	}
	sched.add(ctx, newStep, waits)
	return nil
}

// newScheduler creates new scheduler.
func newScheduler(ctx context.Context, opt schedulerOption) *scheduler {
	if opt.EnableTrace {
		clog.Infof(ctx, "schedule: enable trace")
	}
	return &scheduler{
		path:   opt.Path,
		hashFS: opt.HashFS,
		plan: &plan{
			m: make(map[string]bool),
			// preallocate capacity for performance optimization.
			q:     make(chan *Step, 10000),
			waits: make(map[string][]*Step),
		},
		stats:       &stats{},
		enableTrace: opt.EnableTrace,
	}
}

// mark marks target (exec root relative) as source file.
func (s *scheduler) mark(ctx context.Context, target string) error {
	_, err := s.hashFS.Stat(ctx, s.path.ExecRoot, target)
	if err != nil {
		return fmt.Errorf("missing source %q: %v", target, err)
	}
	s.plan.m[target] = true
	return nil
}

// marked checks target is already marked.
func (s *scheduler) marked(target string) bool {
	return s.plan.m[target]
}

func (s *scheduler) progressReport(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	ui.PrintLines(msg)
}

// finish finishes the scheduling.
func (s *scheduler) finish(ctx context.Context, d time.Duration) {
	s.plan.mu.Lock()
	defer s.plan.mu.Unlock()
	nready := len(s.plan.q) + len(s.plan.ready)
	npendings := s.plan.npendings
	s.progressReport("schedule pending:%d+ready:%d (node:%d edge:%d) %s\n", npendings, nready, len(s.plan.m), s.visited, d)
}

// add adds new stepDef to run.
func (s *scheduler) add(ctx context.Context, stepDef StepDef, waits []string) {
	s.plan.mu.Lock()
	defer s.plan.mu.Unlock()
	defer func() {
		if time.Since(s.lastProgress) < 1*time.Second {
			return
		}
		nready := len(s.plan.q) + len(s.plan.ready)
		npendings := s.plan.npendings
		s.progressReport("schedule pending:%d+ready:%d (node:%d edge:%d)", npendings, nready, len(s.plan.m), s.visited)
		s.lastProgress = time.Now()
	}()
	s.stats.incTotal()
	step := newStep(stepDef, waits)
	if step.ReadyToRun("", "") {
		clog.Infof(ctx, "step state: %s ready to run", step.String())
		select {
		case s.plan.q <- step:
		default:
			step.queueTime = time.Now()
			step.queueSize = len(s.plan.ready)
			s.plan.ready = append(s.plan.ready, step)
		}
		return
	}
	clog.Infof(ctx, "pending to run: %s (waits: %d)", step, step.NumWaits())
	s.plan.npendings++
	for _, w := range waits {
		s.plan.waits[w] = append(s.plan.waits[w], step)
	}
}

type planStats struct {
	npendings int
	nready    int
}

func (p *plan) stats() planStats {
	p.mu.Lock()
	defer p.mu.Unlock()
	return planStats{
		npendings: p.npendings,
		nready:    len(p.q) + len(p.ready),
	}
}

func (p *plan) pushReady() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.ready) == 0 {
		return
	}
	select {
	case p.q <- p.ready[0]:
		p.ready[0].queueDuration = time.Since(p.ready[0].queueTime)
		// Deallocate p.ready[0] explicitly.
		copy(p.ready, p.ready[1:])
		p.ready[len(p.ready)-1] = nil
		p.ready = p.ready[:len(p.ready)-1]
	default:
	}
}

func (p *plan) done(ctx context.Context, step *Step, outs []string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Before processing the completed step,
	// send ready steps from p.ready to p.q and resize p.ready.
	i := 0
	for _, s := range p.ready {
		select {
		case p.q <- s:
			s.queueDuration = time.Since(s.queueTime)
		default:
			p.ready[i] = s
			i++
		}
	}
	for j := i; j < len(p.ready); j++ {
		p.ready[j] = nil
	}
	p.ready = p.ready[:i]

	// Unblock waiting steps and send them to the queue if they are ready.
	npendings := p.npendings
	nready := 0
	ready := make([]*Step, 0, len(outs))
	for _, out := range outs {
		if log.V(1) {
			clog.Infof(ctx, "done %s", out)
		}
		i = 0
		for _, s := range p.waits[out] {
			if s.ReadyToRun(step.String(), out) {
				p.npendings--
				nready++
				if log.V(1) {
					clog.Infof(ctx, "step state: %s ready to run", s.String())
				}
				select {
				case p.q <- s:
				default:
					s.queueTime = time.Now()
					s.queueSize = len(ready)
					ready = append(ready, s)
				}
				continue
			}
			p.waits[out][i] = s
			i++
		}
		for j := i; j < len(p.waits[out]); j++ {
			p.waits[out][j] = nil
		}
		if i == 0 {
			delete(p.waits, out)
			continue
		}
		p.waits[out] = p.waits[out][:i]
	}
	if log.V(1) {
		if nready > 0 {
			clog.Infof(ctx, "trigger %d. pendings %d -> %d", nready, npendings, p.npendings)
		} else {
			clog.Infof(ctx, "zero-trigger outs=%q", outs)
		}
	}
	p.ready = append(p.ready, ready...)
	if len(p.ready) == 0 && p.npendings == 0 && !p.closed {
		p.closed = true
		clog.Infof(ctx, "no step in pending. closing q")
		close(p.q)
	}
}

func (p *plan) dump(ctx context.Context) {
	p.mu.Lock()
	defer p.mu.Unlock()
	clog.Infof(ctx, "queue = %d pendings=%d", len(p.q), p.npendings)
	clog.Infof(ctx, "closed=%t", p.closed)
	var steps []*Step
	seen := make(map[*Step]bool)
	waits := make(map[string]bool)
	ready := make([]string, 0, len(p.ready))
	for _, s := range p.ready {
		ready = append(ready, s.String())
		seen[s] = true
		steps = append(steps, s)
	}
	clog.Infof(ctx, "ready=%q", ready)
	for node, ws := range p.waits {
		waits[node] = true
		for _, s := range ws {
			if seen[s] {
				continue
			}
			seen[s] = true
			steps = append(steps, s)
		}
	}
	for _, s := range steps {
		for _, o := range s.def.Outputs() {
			if !waits[o] {
				clog.Infof(ctx, "step %s output:%s no trigger", s, o)
				continue
			}
			delete(waits, o)
		}
	}
	outs := make([]string, 0, len(waits))
	for k := range waits {
		outs = append(outs, k)
	}
	sort.Strings(outs)
	clog.Infof(ctx, "waits=%d no-trigger=%d", len(p.waits), len(outs))
	clog.Infof(ctx, "no steps will trigger %q", outs)
}
