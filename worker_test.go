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

	want = 123
	worker.handle(func() {
		got = 123
	})

	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("worker.handle should panic")
		}
	}()

	worker.executor.conf.recoverFunc = nil
	worker.handle(func() {
		panic(want)
	})
}

// go test -v -cover -run=^TestWorkerWaitingTasks$
func TestWorkerWaitingTasks(t *testing.T) {
	taskQueue := make(chan Task, 4)
	worker := &worker{taskQueue: taskQueue}

	if worker.WaitingTasks() != len(taskQueue) {
		t.Fatalf("got %d != want %d", worker.WaitingTasks(), len(taskQueue))
	}

	if worker.WaitingTasks() != 0 {
		t.Fatalf("got %d != 0", worker.WaitingTasks())
	}

	taskQueue <- nil
	taskQueue <- nil

	if worker.WaitingTasks() != len(taskQueue) {
		t.Fatalf("got %d != want %d", worker.WaitingTasks(), len(taskQueue))
	}

	if worker.WaitingTasks() != 2 {
		t.Fatalf("got %d != 2", worker.WaitingTasks())
	}
}
