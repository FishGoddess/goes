// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"context"
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestExecutor$
func TestExecutor(t *testing.T) {
	workers := 16
	executor := NewExecutor(uint(workers))
	defer executor.Close()

	var countMap = make(map[int64]int, 16)
	var lock sync.Mutex

	ctx := context.Background()
	totalCount := 10 * workers
	for i := 0; i < totalCount; i++ {
		executor.Submit(ctx, func() {
			now := time.Now().UnixMilli() / 10

			lock.Lock()
			countMap[now] = countMap[now] + 1
			lock.Unlock()

			time.Sleep(10 * time.Millisecond)
		})
	}

	executor.Close()

	gotTotalCount := 0
	for now, count := range countMap {
		gotTotalCount = gotTotalCount + count

		if count != workers {
			t.Fatalf("now %d: count %d != workers %d", now, count, workers)
		}
	}

	if gotTotalCount != totalCount {
		t.Fatalf("gotTotalCount %d != totalCount %d", gotTotalCount, totalCount)
	}
}

// go test -v -cover -run=^TestExecutorMinMax$
func TestExecutorMinMax(t *testing.T) {
	executor := NewExecutor(minWorkers - 1)

	if executor.workers != minWorkers {
		t.Fatalf("got %d != want %d", executor.workers, minWorkers)
	}

	executor.Close()
	executor = NewExecutor(maxWorkers + 1)

	if executor.workers != maxWorkers {
		t.Fatalf("got %d != want %d", executor.workers, maxWorkers)
	}

	executor.Close()
}

// go test -v -cover -run=^TestExecutorContext$
func TestExecutorContext(t *testing.T) {
	executor := NewExecutor(4, WithQueueSize(1))
	defer executor.Close()

	got := uint(cap(executor.tasks))
	if got != executor.conf.queueSize {
		t.Fatalf("got %d != want %d", got, executor.conf.queueSize)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	for i := uint(0); i <= executor.workers; i++ {
		err := executor.Submit(ctx, func() {
			time.Sleep(200 * time.Millisecond)
		})

		if err != nil {
			t.Fatal(err)
		}
	}

	err := executor.Submit(ctx, func() {})
	if err != context.DeadlineExceeded {
		t.Fatalf("got %+v != want %+v", err, context.DeadlineExceeded)
	}
}

// go test -v -cover -run=^TestExecutorClose$
func TestExecutorClose(t *testing.T) {
	executor := NewExecutor(4, WithQueueSize(1))
	defer executor.Close()

	if executor.closed.Load() {
		t.Fatal("executor is closed")
	}

	ctx := context.Background()
	for i := uint(0); i <= executor.workers; i++ {
		err := executor.Submit(ctx, func() {
			time.Sleep(200 * time.Millisecond)
		})

		if err != nil {
			t.Fatal(err)
		}
	}

	go func() {
		err := executor.Submit(ctx, func() {})
		if err != ErrExecutorClosed {
			t.Errorf("got %+v != want %+v", err, ErrExecutorClosed)
		}
	}()

	time.Sleep(10 * time.Millisecond)
	executor.Close()

	if !executor.closed.Load() {
		t.Fatal("executor not closed")
	}

	err := executor.Submit(ctx, func() {})
	if err != ErrExecutorClosed {
		t.Errorf("got %+v != want %+v", err, ErrExecutorClosed)
	}
}
