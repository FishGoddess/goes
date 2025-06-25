// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type roundRobinWorkers struct {
	workers []*worker
	index   int
}

func newRoundRobinWorkers(workerNum int) *roundRobinWorkers {
	rrWorkers := &roundRobinWorkers{
		workers: make([]*worker, 0, workerNum),
		index:   -1,
	}

	return rrWorkers
}

// Add adds a worker to workers.
func (rrw *roundRobinWorkers) Add(worker *worker) {
	rrw.workers = append(rrw.workers, worker)
}

// Next returns the next worker from workers.
func (rrw *roundRobinWorkers) Next() *worker {
	if len(rrw.workers) <= 0 {
		return nil
	}

	if rrw.index++; rrw.index >= len(rrw.workers) {
		rrw.index = 0
	}

	worker := rrw.workers[rrw.index]
	return worker
}

// Done will call done method on all workers.
func (rrw *roundRobinWorkers) Done() {
	for _, worker := range rrw.workers {
		worker.Done()
	}
}
