// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"

	"github.com/FishGoddess/goes/pkg/spinlock"
)

type Pool struct {
	workers []*worker
	index   int
	closed  bool

	wg   sync.WaitGroup
	lock sync.Locker
}

func NewPool(size int, workerLimit int) *Pool {
	if size <= 0 {
		panic("goes: pool size <= 0")
	}

	pool := &Pool{
		workers: make([]*worker, 0, size),
		index:   0,
		closed:  false,
		lock:    spinlock.New(),
	}

	for range size {
		worker := newWorker(pool, workerLimit)
		pool.workers = append(pool.workers, worker)
	}

	return pool
}

func (p *Pool) Go(f func()) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.closed {
		return
	}

	worker := p.workers[p.index]
	worker.Accept(f)

	p.index++

	if p.index >= len(p.workers) {
		p.index = 0
	}
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, worker := range p.workers {
		worker.Done()
	}

	p.wg.Wait()
	p.closed = true
}
