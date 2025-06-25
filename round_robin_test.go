// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "testing"

// go test -v -cover -run=^TestRoundRobinWorkers$
func TestRoundRobinWorkers(t *testing.T) {
	workerNum := 16
	rrWorkers := newRoundRobinWorkers(workerNum)

	testWorkers := make([]*worker, workerNum)
	for _, worker := range testWorkers {
		rrWorkers.Add(worker)
	}

	if len(rrWorkers.workers) != len(testWorkers) {
		t.Fatalf("len(rrWorkers.workers) %d != len(testWorkers) %d", len(rrWorkers.workers), len(testWorkers))
	}

	for i, worker := range testWorkers {
		got := rrWorkers.workers[i]
		if got != worker {
			t.Fatalf("got %p != worker %p", got, worker)
		}
	}

	for _, worker := range testWorkers {
		gotNext := rrWorkers.Next()
		if gotNext != worker {
			t.Fatalf("gotNext %p != worker %p", gotNext, worker)
		}
	}
}
