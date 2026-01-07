// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

const (
	minWorkers = 1
	maxWorkers = 10000
)

var (
	ErrExecutorClosed = errors.New("goes: executor is closed")
)

// Executor starts some workers to do tasks concurrently.
type Executor struct {
	conf *config

	workers uint
	tasks   chan Task
	done    chan struct{}
	closed  atomic.Bool
	group   sync.WaitGroup
}

// NewExecutor creates a executor with workers.
func NewExecutor(workers uint, opts ...Option) *Executor {
	conf := newConfig().apply(opts...)

	if workers < minWorkers {
		workers = minWorkers
	}

	if workers > maxWorkers {
		workers = maxWorkers
	}

	executor := &Executor{
		conf:    conf,
		workers: workers,
		tasks:   make(chan Task, conf.queueSize),
		done:    make(chan struct{}),
	}

	for range workers {
		executor.group.Go(executor.worker)
	}

	return executor
}

func (e *Executor) worker() {
	for task := range e.tasks {
		task.Do(e.conf.recovery)
	}
}

// Submit submits a task to executor and returns an error if failed.
func (e *Executor) Submit(ctx context.Context, task Task) error {
	if e.closed.Load() {
		return ErrExecutorClosed
	}

	select {
	case e.tasks <- task:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-e.done:
		return ErrExecutorClosed
	}
}

// Close closes the executor and waits all tasks to be done.
func (e *Executor) Close() {
	if !e.closed.CompareAndSwap(false, true) {
		return
	}

	close(e.done)
	close(e.tasks)

	e.group.Wait()
}
