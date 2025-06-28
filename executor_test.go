// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestExecutor$
func TestExecutor(t *testing.T) {
	workerNum := 16
	executor := NewExecutor(workerNum)

	var countMap = make(map[int64]int, 16)
	var lock sync.Mutex

	totalCount := 10 * workerNum
	for i := 0; i < totalCount; i++ {
		executor.Submit(func() {
			now := time.Now().UnixMilli() / 10

			lock.Lock()
			countMap[now] = countMap[now] + 1
			lock.Unlock()

			time.Sleep(10 * time.Millisecond)
		})
	}

	executor.Close()
	executor.Wait()

	gotTotalCount := 0
	for now, count := range countMap {
		gotTotalCount = gotTotalCount + count

		if count != workerNum {
			t.Fatalf("now %d: count %d != workerNum %d", now, count, workerNum)
		}
	}

	if gotTotalCount != totalCount {
		t.Fatalf("gotTotalCount %d != totalCount %d", gotTotalCount, totalCount)
	}
}

// go test -v -cover -run=^TestExecutorError$
func TestExecutorError(t *testing.T) {
	workerNum := 16
	executor := NewExecutor(workerNum)
	executor.Close()

	err := executor.Submit(func() {})
	if err != ErrExecutorIsClosed {
		t.Fatalf("err %v != ErrExecutorIsClosed %v", err, ErrExecutorIsClosed)
	}

	executor = NewExecutor(workerNum)
	executor.scheduler = &roundRobinScheduler{}

	err = executor.Submit(func() {})
	if err != ErrWorkerIsNil {
		t.Fatalf("err %v != ErrWorkerIsNil %v", err, ErrWorkerIsNil)
	}
}

// go test -v -cover -run=^TestExecutorAvailableWorkers$
func TestExecutorAvailableWorkers(t *testing.T) {
	workerNum := 16
	executor := NewExecutor(workerNum)

	if len(executor.workers) != 1 {
		t.Fatalf("len(executor.workers) %d != 1", len(executor.workers))
	}

	if executor.AvailableWorkers() != 1 {
		t.Fatalf("executor.AvailableWorkers() %d != 1", executor.AvailableWorkers())
	}

	executor.workers = make([]*worker, workerNum)

	if executor.AvailableWorkers() != workerNum {
		t.Fatalf("executor.AvailableWorkers() %d != workerNum %d", executor.AvailableWorkers(), workerNum)
	}
}

// go test -v -cover -run=^TestExecutorSpawnWorker$
func TestExecutorSpawnWorker(t *testing.T) {
	workerNum := 16
	executor := NewExecutor(workerNum)

	if executor.AvailableWorkers() != 1 {
		t.Fatalf("executor.AvailableWorkers() %d != 1", executor.AvailableWorkers())
	}

	for range workerNum {
		executor.Submit(func() {})
		time.Sleep(time.Millisecond)
	}

	if executor.AvailableWorkers() != 1 {
		t.Fatalf("executor.AvailableWorkers() %d != 1", executor.AvailableWorkers())
	}

	for range workerNum * 2 {
		executor.Submit(func() {
			time.Sleep(time.Millisecond)
		})
	}

	if executor.AvailableWorkers() != workerNum {
		t.Fatalf("executor.AvailableWorkers() %d != workerNum %d", executor.AvailableWorkers(), workerNum)
	}
}
