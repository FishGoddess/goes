// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type roundRobinScheduler struct {
	workers []*worker
	index   int
}

func newRoundRobinScheduler(workers []*worker) *roundRobinScheduler {
	scheduler := &roundRobinScheduler{
		workers: workers,
		index:   -1,
	}

	return scheduler
}

// Set sets the workers to scheduler.
func (rrs *roundRobinScheduler) Set(workers []*worker) {
	rrs.workers = workers
}

// Get gets a worker from scheduler.
func (rrs *roundRobinScheduler) Get() *worker {
	if len(rrs.workers) <= 0 {
		return nil
	}

	if rrs.index++; rrs.index >= len(rrs.workers) {
		rrs.index = 0
	}

	worker := rrs.workers[rrs.index]
	return worker
}
