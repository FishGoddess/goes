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

// go test -v -cover -run=^TestExecutorWorkerNum$
func TestExecutorWorkerNum(t *testing.T) {
	workerNum := 16
	workers := make([]*worker, workerNum)

	executor := NewExecutor(workerNum)
	executor.workers = workers

	if executor.WorkerNum() != workerNum {
		t.Fatalf("executor.WorkerNum() %d != workerNum %d", executor.WorkerNum(), workerNum)
	}
}
