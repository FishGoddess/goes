// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "testing"

// go test -v -cover -run=^TestRandomWorkers$
func TestRandomWorkers(t *testing.T) {
	workerNum := 16
	rWorkers := newRandomWorkers(workerNum)

	testWorkers := make([]*worker, 0, workerNum)
	for i := 0; i < workerNum; i++ {
		worker := new(worker)

		testWorkers = append(testWorkers, worker)
		rWorkers.Add(worker)
	}

	if len(rWorkers.workers) != len(testWorkers) {
		t.Fatalf("len(rWorkers.workers) %d != len(testWorkers) %d", len(rWorkers.workers), len(testWorkers))
	}

	for i, worker := range testWorkers {
		got := rWorkers.workers[i]
		if got != worker {
			t.Fatalf("got %p != worker %p", got, worker)
		}
	}

	for range testWorkers {
		gotNext := rWorkers.Next()
		if gotNext == nil {
			t.Fatal("gotNext is nil")
		}
	}
}
