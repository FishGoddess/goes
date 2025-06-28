// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"errors"
	"sync"
)

var (
	ErrExecutorIsClosed = errors.New("goes: executor is closed")
	ErrWorkerIsNil      = errors.New("goes: worker is nil")
)

// Task is a function can be executed by executor.
type Task = func()

// Executor executes tasks concurrently using limited goroutines.
// You can specify the number of workers and the queue size of each worker.
type Executor struct {
	conf *config

	workers   []*worker
	scheduler scheduler
	closed    bool

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

	workers := make([]*worker, 0, conf.workerNum)
	executor := &Executor{
		conf:      conf,
		workers:   workers,
		scheduler: conf.newScheduler(),
		closed:    false,
		lock:      conf.newLocker(),
	}

	executor.spawnWorker()
	return executor
}

func (e *Executor) spawnWorker() *worker {
	worker := newWorker(e)

	e.workers = append(e.workers, worker)
	e.scheduler.Set(e.workers)
	return worker
}

// AvailableWorkers returns the number of workers available.
func (e *Executor) AvailableWorkers() int {
	e.lock.Lock()
	defer e.lock.Unlock()

	return len(e.workers)
}

// Submit submits a task to be handled by workers.
func (e *Executor) Submit(task Task) error {
	e.lock.Lock()

	if e.closed {
		e.lock.Unlock()

		return ErrExecutorIsClosed
	}

	worker := e.scheduler.Get()
	if worker == nil {
		e.lock.Unlock()

		return ErrWorkerIsNil
	}

	// 1. We don't need to create a new worker if we got a worker with no tasks.
	// 2. The number of workers has reached the limit, so we can only use the worker we got.
	if worker.WaitingTasks() <= 0 || len(e.workers) >= e.conf.workerNum {
		e.lock.Unlock()

		worker.Accept(task)
		return nil
	}

	worker = e.spawnWorker()
	e.lock.Unlock()

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

	for _, worker := range e.workers {
		worker.Done()
	}

	e.Wait()
	e.closed = true
	e.workers = nil
	e.scheduler.Set(nil)
}
