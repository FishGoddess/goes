// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"testing"
)

// go test -v -cover -run=^TestRoundRobinScheduler$
func TestRoundRobinScheduler(t *testing.T) {
	workerNum := 16
	workers := make([]*worker, 0, workerNum)
	for range workerNum {
		workers = append(workers, new(worker))
	}

	scheduler := newRoundRobinScheduler(workers...)
	if fmt.Sprintf("%p", scheduler.workers) != fmt.Sprintf("%p", workers) {
		t.Fatalf("scheduler.workers %p != workers %p", scheduler.workers, workers)
	}

	got := len(scheduler.workers)
	want := len(workers)
	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}

	scheduler.Set(workers)
	if fmt.Sprintf("%p", scheduler.workers) != fmt.Sprintf("%p", workers) {
		t.Fatalf("scheduler.workers %p != workers %p", scheduler.workers, workers)
	}

	got = len(scheduler.workers)
	want = len(workers)
	if got != want {
		t.Fatalf("got %d != want %d", got, want)
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
