// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"errors"
	"sync"
)

var (
	ErrExecutorClosed = errors.New("goes: executor is closed")
	ErrWorkerIsNil    = errors.New("goes: worker is nil")
)

// Task is a function can be executed by executor.
type Task = func()

// Executor executes tasks concurrently using limited goroutines.
// You can specify the number of workers and the queue size of each worker.
type Executor struct {
	conf *config

	workers workers
	closed  bool

	wg   sync.WaitGroup
	lock sync.Locker
}

// NewExecutor creates a new executor with given worker number and options.
func NewExecutor(workerNum int, opts ...Option) *Executor {
	conf := newDefaultConfig(workerNum)
	for _, opt := range opts {
		opt.applyTo(conf)
	}

	if conf.workerNum <= 0 {
		panic("goes: executor's worker num <= 0")
	}

	if conf.workerQueueSize <= 0 {
		panic("goes: worker's queue size <= 0")
	}

	executor := &Executor{
		conf:    conf,
		workers: newRoundRobinWorkers(conf.workerNum),
		closed:  false,
		lock:    conf.newLocker(),
	}

	for range conf.workerNum {
		worker := newWorker(executor)
		executor.workers.Add(worker)
	}

	return executor
}

// Submit submits a task to be handled by workers.
func (e *Executor) Submit(task Task) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.closed {
		return ErrExecutorClosed
	}

	worker := e.workers.Next()
	if worker == nil {
		return ErrWorkerIsNil
	}

	worker.Accept(task)
	return nil
}

// Wait waits all tasks to be handled.
func (e *Executor) Wait() {
	e.wg.Wait()
}

// Close closes the executor after handling all tasks.
func (e *Executor) Close() {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.workers.Done()
	e.wg.Wait()
	e.closed = true
}
