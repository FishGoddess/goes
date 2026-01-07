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

var ErrExecutorClosed = errors.New("goes: executor is closed")

type Executor struct {
	conf *config

	tasks  chan Task
	done   chan struct{}
	closed atomic.Bool

	group sync.WaitGroup
}

func NewExecutor(workers uint, opts ...Option) *Executor {
	conf := newConfig().apply(opts...)

	if workers < 1 {
		panic("goes: workers < 1")
	}

	executor := &Executor{
		conf:  conf,
		tasks: make(chan Task, conf.queueSize),
		done:  make(chan struct{}),
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

func (e *Executor) Close() error {
	if !e.closed.CompareAndSwap(false, true) {
		return nil
	}

	close(e.done)
	close(e.tasks)

	e.group.Wait()
	return nil
}
