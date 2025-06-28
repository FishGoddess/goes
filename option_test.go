// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"sync"
	"testing"

	"github.com/FishGoddess/goes/pkg/spinlock"
)

// go test -v -cover -run=^TestWithWorkerQueueSize$
func TestWithWorkerQueueSize(t *testing.T) {
	workerNum := 16
	workerQueueSize := 256
	conf := newDefaultConfig(workerNum)
	WithWorkerQueueSize(workerQueueSize)(conf)

	if conf.workerQueueSize != workerQueueSize {
		t.Fatalf("conf.workerQueueSize %d != workerQueueSize %d", conf.workerQueueSize, workerQueueSize)
	}
}

// go test -v -cover -run=^TestWithRecoverFunc$
func TestWithRecoverFunc(t *testing.T) {
	workerNum := 16
	recoverFunc := func(r any) {}
	conf := newDefaultConfig(workerNum)
	WithRecoverFunc(recoverFunc)(conf)

	if fmt.Sprintf("%p", conf.recoverFunc) != fmt.Sprintf("%p", recoverFunc) {
		t.Fatalf("conf.recoverFunc %p != recoverFunc %p", conf.recoverFunc, recoverFunc)
	}
}

// go test -v -cover -run=^TestWithNewLockerFunc$
func TestWithNewLockerFunc(t *testing.T) {
	workerNum := 16
	newLockerFunc := func() sync.Locker { return nil }
	conf := newDefaultConfig(workerNum)
	WithNewLockerFunc(newLockerFunc)(conf)

	if fmt.Sprintf("%p", conf.newLockerFunc) != fmt.Sprintf("%p", newLockerFunc) {
		t.Fatalf("conf.newLockerFunc %p != newLockerFunc %p", conf.newLockerFunc, newLockerFunc)
	}
}

// go test -v -cover -run=^TestWithSpinLock$
func TestWithSpinLock(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	WithSpinLock()(conf)

	got := conf.newLockerFunc()
	if _, ok := got.(*spinlock.Lock); !ok {
		t.Fatalf("got %T is not *spinlock.Lock", got)
	}
}

// go test -v -cover -run=^TestWithSyncMutex$
func TestWithSyncMutex(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	WithSyncMutex()(conf)

	got := conf.newLockerFunc()
	if _, ok := got.(*sync.Mutex); !ok {
		t.Fatalf("got %T is not *sync.Mutex", got)
	}
}

// go test -v -cover -run=^TestWithRoundRobinScheduler$
func TestWithRoundRobinScheduler(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	WithRoundRobinScheduler()(conf)

	workers := make([]*worker, workerNum)
	got := conf.newSchedulerFunc(workers)

	scheduler, ok := got.(*roundRobinScheduler)
	if !ok {
		t.Fatalf("got %T is not *roundRobinScheduler", got)
	}

	if cap(scheduler.workers) != workerNum {
		t.Fatalf("cap(scheduler.workers) %d != workerNum %d", cap(scheduler.workers), workerNum)
	}
}

// go test -v -cover -run=^TestWithRandomScheduler$
func TestWithRandomScheduler(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	WithRandomScheduler()(conf)

	workers := make([]*worker, workerNum)
	got := conf.newSchedulerFunc(workers)

	scheduler, ok := got.(*randomScheduler)
	if !ok {
		t.Fatalf("got %T is not *randomScheduler", got)
	}

	if cap(scheduler.workers) != workerNum {
		t.Fatalf("cap(scheduler.workers) %d != workerNum %d", cap(scheduler.workers), workerNum)
	}
}
