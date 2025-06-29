// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestNewExecutor$
func TestNewExecutor(t *testing.T) {
	workerNum := 16

	t.Run("ok", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("r should not panic")
			}
		}()

		executor := NewExecutor(workerNum, WithWorkerQueueSize(256), WithPurgeActive(time.Minute, time.Minute))
		defer executor.Close()
	})

	testCase := func(t *testing.T, workerNum int, opts ...Option) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("r should panic")
			}
		}()

		executor := NewExecutor(workerNum, opts...)
		defer executor.Close()
	}

	t.Run("worker num", func(t *testing.T) {
		testCase(t, 0)
	})

	t.Run("worker queue size", func(t *testing.T) {
		testCase(t, workerNum, WithWorkerQueueSize(0))
	})

	t.Run("purge task 1", func(t *testing.T) {
		testCase(t, workerNum, WithPurgeActive(time.Millisecond, 0))
	})

	t.Run("purge task 2", func(t *testing.T) {
		testCase(t, workerNum, WithPurgeActive(0, time.Millisecond))
	})

	t.Run("purge task 3", func(t *testing.T) {
		testCase(t, workerNum, WithPurgeActive(time.Millisecond, time.Millisecond))
	})
}

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

	got := executor.AvailableWorkers()
	if got != 1 {
		t.Fatalf("got %d != 1", got)
	}

	executor.workers = make([]*worker, workerNum)

	got = executor.AvailableWorkers()
	if got != workerNum {
		t.Fatalf("got %d != workerNum %d", got, workerNum)
	}
}

// go test -v -cover -run=^TestExecutorSpawnWorker$
func TestExecutorSpawnWorker(t *testing.T) {
	workerNum := 16
	executor := NewExecutor(workerNum)

	got := executor.AvailableWorkers()
	if got != 1 {
		t.Fatalf("got %d != 1", got)
	}

	for range workerNum {
		executor.Submit(func() {})
		time.Sleep(time.Millisecond)
	}

	got = executor.AvailableWorkers()
	if got != 1 {
		t.Fatalf("got %d != 1", got)
	}

	for range workerNum * 2 {
		executor.Submit(func() {
			time.Sleep(time.Millisecond)
		})
	}

	got = executor.AvailableWorkers()
	if got != workerNum {
		t.Fatalf("got %d != workerNum %d", got, workerNum)
	}
}

// go test -v -cover -run=^TestExecutorDynamicScaling$
func TestExecutorDynamicScaling(t *testing.T) {
	testCase := func(t *testing.T, workerNum int) {
		executor := NewExecutor(workerNum)
		executor.conf.purgeInterval = time.Millisecond
		executor.conf.workerLifetime = 2 * time.Millisecond
		executor.runPurgeTask()
		defer executor.Close()

		got := executor.AvailableWorkers()
		if got != 1 {
			t.Fatalf("got %d != 1", got)
		}

		for range workerNum * 2 {
			executor.Submit(func() {
				time.Sleep(time.Millisecond)
			})
		}

		got = executor.AvailableWorkers()
		if got != workerNum {
			t.Fatalf("got %d != workerNum %d", got, workerNum)
		}

		time.Sleep(500 * time.Microsecond)

		got = executor.AvailableWorkers()
		if got != workerNum {
			t.Fatalf("got %d != workerNum %d", got, workerNum)
		}

		time.Sleep(5 * time.Millisecond)

		got = executor.AvailableWorkers()
		if got != 1 {
			t.Fatalf("got %d != workerNum %d", got, workerNum)
		}

		for range workerNum * 2 {
			executor.Submit(func() {
				time.Sleep(time.Millisecond)
			})
		}

		got = executor.AvailableWorkers()
		if got != workerNum {
			t.Fatalf("got %d != workerNum %d", got, workerNum)
		}

		time.Sleep(5 * time.Millisecond)

		got = executor.AvailableWorkers()
		if got != 1 {
			t.Fatalf("got %d != workerNum %d", got, workerNum)
		}
	}

	workerNums := []int{1, 2, 4, 16, 64, 256, 1024}
	for _, workerNum := range workerNums {
		name := fmt.Sprintf("worker num %d", workerNum)
		t.Run(name, func(t *testing.T) {
			testCase(t, workerNum)
		})
	}
}
