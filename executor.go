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

	executor := &Executor{
		conf:   conf,
		closed: false,
		lock:   conf.newLocker(),
	}

	workers := make([]*worker, 0, conf.workerNum)
	for range conf.workerNum {
		worker := newWorker(executor)
		workers = append(workers, worker)
	}

	executor.workers = workers
	executor.scheduler = conf.newScheduler(workers)
	return executor
}

// WorkerNum returns the number of workers in the executor.
func (e *Executor) WorkerNum() int {
	e.lock.Lock()
	defer e.lock.Unlock()

	return len(e.workers)
}

// Submit submits a task to be handled by workers.
func (e *Executor) Submit(task Task) error {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.closed {
		return ErrExecutorIsClosed
	}

	worker := e.scheduler.Get()
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

	for _, worker := range e.workers {
		worker.Done()
	}

	e.wg.Wait()
	e.closed = true
}
