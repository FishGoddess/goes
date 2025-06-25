// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "testing"

// go test -v -cover -run=^TestWorkerHandle$
func TestWorkerHandle(t *testing.T) {
	got := 0
	want := 666

	executor := &Executor{
		conf: &config{
			recoverFunc: func(r any) {
				got = r.(int)
			},
		},
	}

	worker := &worker{executor: executor}
	worker.handle(func() {
		panic(want)
	})

	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}

	worker.handle(func() {
		got = 123
		want = 123
	})

	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}
}
