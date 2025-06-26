// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"math/rand"
	"time"
)

type randomWorkers struct {
	workers []*worker
	random  *rand.Rand
}

func newRandomWorkers(workerNum int) *randomWorkers {
	rWorkers := &randomWorkers{
		workers: make([]*worker, 0, workerNum),
		random:  rand.New(rand.NewSource(time.Now().Unix())),
	}

	return rWorkers
}

// Add adds a worker to workers.
func (rw *randomWorkers) Add(worker *worker) {
	rw.workers = append(rw.workers, worker)
}

// Next returns the next worker from workers.
func (rw *randomWorkers) Next() *worker {
	if len(rw.workers) <= 0 {
		return nil
	}

	index := rw.random.Intn(len(rw.workers))
	worker := rw.workers[index]
	return worker
}

// Done will call done method on all workers.
func (rw *randomWorkers) Done() {
	for _, worker := range rw.workers {
		worker.Done()
	}
}
