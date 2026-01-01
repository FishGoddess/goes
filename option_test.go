// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"fmt"
	"testing"
)

// go test -v -cover -run=^TestWithQueueSize$
func TestWithQueueSize(t *testing.T) {
	conf := &config{queueSize: 0}

	queueSize := uint(1024)
	WithQueueSize(queueSize)(conf)

	if conf.queueSize != queueSize {
		t.Fatalf("got %d != want %d", conf.queueSize, queueSize)
	}
}

// go test -v -cover -run=^WithRecovery$
func TestWithRecovery(t *testing.T) {
	conf := &config{recovery: nil}

	recovery := func(r any) {}
	WithRecovery(recovery)(conf)

	got := fmt.Sprintf("%p", conf.recovery)
	want := fmt.Sprintf("%p", recovery)
	if got != want {
		t.Fatalf("got %s != want %s", got, want)
	}
}
