// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "testing"

// go test -v -cover -run=^TestRoundRobinScheduler$
func TestRoundRobinScheduler(t *testing.T) {
	workerNum := 16
	workers := make([]*worker, 0, workerNum)
	for i := 0; i < workerNum; i++ {
		worker := new(worker)
		workers = append(workers, worker)
	}

	scheduler := newRoundRobinScheduler(workers)
	scheduler.Set(workers)

	if len(scheduler.workers) != len(workers) {
		t.Fatalf("len(scheduler.workers) %d != len(workers) %d", len(scheduler.workers), len(workers))
	}

	for i, worker := range workers {
		got := scheduler.workers[i]
		if got != worker {
			t.Fatalf("got %p != worker %p", got, worker)
		}
	}

	for _, worker := range workers {
		gotNext := scheduler.Get()
		if gotNext != worker {
			t.Fatalf("gotNext %p != worker %p", gotNext, worker)
		}
	}
}
