// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -cover -run=^TestWithWorkerQueueSize$
func TestWithWorkerQueueSize(t *testing.T) {
	conf := &config{workerQueueSize: 0}

	workerQueueSize := 1024
	WithWorkerQueueSize(workerQueueSize)(conf)

	if conf.workerQueueSize != workerQueueSize {
		t.Fatalf("got %d != want %d", conf.workerQueueSize, workerQueueSize)
	}
}

// go test -v -cover -run=^TestWithNow$
func TestWithNow(t *testing.T) {
	conf := &config{now: nil}

	now := func() time.Time { return time.Now() }
	WithNow(now)(conf)

	got := fmt.Sprintf("%p", conf.now)
	want := fmt.Sprintf("%p", now)
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}

// go test -v -cover -run=^TestWithHandlePanic$
func TestWithHandlePanic(t *testing.T) {
	conf := &config{handlePanic: nil}

	handlePanic := func(r any) {}
	WithHandlePanic(handlePanic)(conf)

	got := fmt.Sprintf("%p", conf.handlePanic)
	want := fmt.Sprintf("%p", handlePanic)
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
