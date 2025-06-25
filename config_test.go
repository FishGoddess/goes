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

// go test -v -cover -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	if conf.workerNum != workerNum {
		t.Fatalf("conf.workerNum %d != workerNum %d", conf.workerNum, workerNum)
	}
}

// go test -v -cover -run=^TestConfigRecover$
func TestConfigRecover(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	conf.recover(0)

	got := 0
	conf.recoverFunc = func(r any) {
		got = r.(int)
	}

	want := 1
	conf.recover(want)

	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}
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

// go test -v -cover -run=^TestConfigNewWorkers$
func TestConfigNewWorkers(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)

	got := conf.newWorkers()
	if _, ok := got.(*roundRobinWorkers); !ok {
		t.Fatalf("got %T is not *roundRobinWorkers", got)
	}

	want := &roundRobinWorkers{}
	conf.newWorkersFunc = func() workers {
		return want
	}

	got = conf.newWorkers()

	if fmt.Sprintf("%p", got) != fmt.Sprintf("%p", want) {
		t.Fatalf("got %p != want %p", got, want)
	}
}
