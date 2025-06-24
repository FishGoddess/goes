// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"
)

type Executor struct {
	conf *config

	workers []*worker
	index   int
	closed  bool

	wg   sync.WaitGroup
	lock sync.Locker
}

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
		workers: make([]*worker, 0, conf.workerNum),
		index:   0,
		closed:  false,
		lock:    conf.newLocker(),
	}

	for range conf.workerNum {
		worker := newWorker(executor)
		executor.workers = append(executor.workers, worker)
	}

	return executor
}

func (e *Executor) Submit(task func()) {
	e.lock.Lock()
	defer e.lock.Unlock()

	if e.closed {
		return
	}

	worker := e.workers[e.index]
	worker.Accept(task)

	if e.index++; e.index >= len(e.workers) {
		e.index = 0
	}
}

func (e *Executor) Wait() {
	e.wg.Wait()
}

func (e *Executor) Close() {
	e.lock.Lock()
	defer e.lock.Unlock()

	for _, worker := range e.workers {
		worker.Done()
	}

	e.wg.Wait()
	e.closed = true
}
