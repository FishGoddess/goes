// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"testing"
	"time"
)

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

	got := worker.WaitingTasks()
	want := len(taskQueue)
	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}

	got = worker.WaitingTasks()
	if got != 0 {
		t.Fatalf("got %d != 0", got)
	}

	taskQueue <- nil
	taskQueue <- nil

	got = worker.WaitingTasks()
	want = len(taskQueue)
	if got != want {
		t.Fatalf("got %d != want %d", got, want)
	}

	got = worker.WaitingTasks()
	if got != 2 {
		t.Fatalf("got %d != 2", got)
	}
}

// go test -v -cover -run=^TestWorkerAcceptTime$
func TestWorkerAcceptTime(t *testing.T) {
	acceptTime := time.Now()
	worker := &worker{acceptTime: acceptTime}

	got := worker.AcceptTime()
	if got != acceptTime {
		t.Fatalf("got %v != acceptTime %v", got, acceptTime)
	}

	acceptTime = time.Unix(123456789, 0)
	worker.SetAcceptTime(acceptTime)

	got = worker.acceptTime
	if got != acceptTime {
		t.Fatalf("got %v != acceptTime %v", got, acceptTime)
	}

	got = worker.AcceptTime()
	if got != acceptTime {
		t.Fatalf("got %v != acceptTime %v", got, acceptTime)
	}
}
