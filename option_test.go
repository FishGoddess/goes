// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"sync"
	"testing"
	"time"

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

// go test -v -cover -run=^TestWithNowFunc$
func TestWithNowFunc(t *testing.T) {
	workerNum := 16
	nowFunc := func() time.Time { return time.Now() }
	conf := newDefaultConfig(workerNum)
	WithNowFunc(nowFunc)(conf)

	if fmt.Sprintf("%p", conf.nowFunc) != fmt.Sprintf("%p", nowFunc) {
		t.Fatalf("conf.nowFunc %p != nowFunc %p", conf.nowFunc, nowFunc)
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

	workers := make([]*worker, 0, workerNum)
	got := conf.newSchedulerFunc(workers...)

	scheduler, ok := got.(*roundRobinScheduler)
	if !ok {
		t.Fatalf("got %T is not *roundRobinScheduler", got)
	}

	gotCap := cap(scheduler.workers)
	if gotCap != workerNum {
		t.Fatalf("gotCap %d != workerNum %d", gotCap, workerNum)
	}
}

// go test -v -cover -run=^TestWithRandomScheduler$
func TestWithRandomScheduler(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	WithRandomScheduler()(conf)

	workers := make([]*worker, 0, workerNum)
	got := conf.newSchedulerFunc(workers...)

	scheduler, ok := got.(*randomScheduler)
	if !ok {
		t.Fatalf("got %T is not *randomScheduler", got)
	}

	gotCap := cap(scheduler.workers)
	if gotCap != workerNum {
		t.Fatalf("gotCap %d != workerNum %d", gotCap, workerNum)
	}
}
