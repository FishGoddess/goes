// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"sync"
	"testing"
)

// go test -v -cover -run=^TestWithQueueSize$
func TestWithQueueSize(t *testing.T) {
	size := 16
	queueSize := 256
	conf := newDefaultConfig(size)
	WithQueueSize(queueSize)(conf)

	if conf.queueSize != queueSize {
		t.Fatalf("conf.queueSize %d != queueSize %d", conf.queueSize, queueSize)
	}
}

// go test -v -cover -run=^TestWithRecoverFunc$
func TestWithRecoverFunc(t *testing.T) {
	size := 16
	recoverFunc := func(r any) {}
	conf := newDefaultConfig(size)
	WithRecoverFunc(recoverFunc)(conf)

	if fmt.Sprintf("%p", conf.recoverFunc) != fmt.Sprintf("%p", recoverFunc) {
		t.Fatalf("conf.recoverFunc %p != recoverFunc %p", conf.recoverFunc, recoverFunc)
	}
}

// go test -v -cover -run=^TestWithNewLockerFunc$
func TestWithNewLockerFunc(t *testing.T) {
	size := 16
	newLockerFunc := func() sync.Locker { return nil }
	conf := newDefaultConfig(size)
	WithNewLockerFunc(newLockerFunc)(conf)

	if fmt.Sprintf("%p", conf.newLockerFunc) != fmt.Sprintf("%p", newLockerFunc) {
		t.Fatalf("conf.newLockerFunc %p != newLockerFunc %p", conf.newLockerFunc, newLockerFunc)
	}
}
