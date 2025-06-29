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

// go test -v -cover -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	if conf.workerNum != workerNum {
		t.Fatalf("conf.workerNum %d != workerNum %d", conf.workerNum, workerNum)
	}
}

// go test -v -cover -run=^TestConfigNow$
func TestConfigNow(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	got := conf.now().Unix()
	want := time.Now().Unix()
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}

	wantTime := time.Unix(123456789, 0)
	conf.nowFunc = func() time.Time {
		return wantTime
	}

	got = conf.now().Unix()
	want = wantTime.Unix()
	if got != want {
		t.Fatalf("got %v != want %v", got, want)
	}
}

// go test -v -cover -run=^TestConfigRecover$
func TestConfigRecover(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	if conf.recoverable() {
		t.Fatalf("conf.recoverable() is wrong")
	}

	got := 0
	conf.recoverFunc = func(r any) {
		got = r.(int)
	}

	if !conf.recoverable() {
		t.Fatalf("conf.recoverable() is wrong")
	}

	want := 1
	conf.recover(want)

	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("conf.recover should panic")
		}
	}()

	conf.recoverFunc = nil
	conf.recover(0)
}

// go test -v -cover -run=^TestConfigNewLocker$
func TestConfigNewLocker(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	got := conf.newLocker()
	if _, ok := got.(*spinlock.Lock); !ok {
		t.Fatalf("got %T is not *spinlock.Lock", got)
	}

	want := &sync.Mutex{}
	conf.newLockerFunc = func() sync.Locker {
		return want
	}

	got = conf.newLocker()
	if fmt.Sprintf("%p", got) != fmt.Sprintf("%p", want) {
		t.Fatalf("got %p != want %p", got, want)
	}
}

// go test -v -cover -run=^TestConfigNewScheduler$
func TestConfigNewScheduler(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	workers := make([]*worker, workerNum)
	got := conf.newScheduler(workers...)

	if _, ok := got.(*roundRobinScheduler); !ok {
		t.Fatalf("got %T is not *roundRobinScheduler", got)
	}

	want := &roundRobinScheduler{}
	conf.newSchedulerFunc = func(workers ...*worker) scheduler {
		want.workers = workers
		return want
	}

	got = conf.newScheduler(workers...)
	if fmt.Sprintf("%p", got) != fmt.Sprintf("%p", want) {
		t.Fatalf("got %p != want %p", got, want)
	}

	scheduler, ok := got.(*roundRobinScheduler)
	if !ok {
		t.Fatalf("got %T is not *roundRobinScheduler", got)
	}

	gotLen := len(scheduler.workers)
	if gotLen != workerNum {
		t.Fatalf("gotLen %d != workerNum %d", gotLen, workerNum)
	}
}
