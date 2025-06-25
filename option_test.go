// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"sync"
	"testing"
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

// go test -v -cover -run=^TestWithSyncMutex$
func TestWithSyncMutex(t *testing.T) {
	workerNum := 16
	conf := newDefaultConfig(workerNum)
	WithSyncMutex()(conf)

	lock := conf.newLockerFunc()
	if _, ok := lock.(*sync.Mutex); !ok {
		t.Fatalf("lock %T is not *sync.Mutex", lock)
	}
}
