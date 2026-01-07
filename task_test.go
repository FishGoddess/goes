// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "testing"

// go test -v -cover -run=^TestTask$
func TestTask(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal(r)
		}
	}()

	var done bool = false
	var task Task = func() { done = true }
	task.Do(nil)

	if !done {
		t.Fatalf("%+v is wrong", done)
	}
}

// go test -v -cover -run=^TestTaskNil$
func TestTaskNil(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal(r)
		}
	}()

	var task Task = nil
	task.Do(nil)
}

// go test -v -cover -run=^TestTaskPanic$
func TestTaskPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("task should panic")
		}
	}()

	var task Task = func() { panic("wow") }
	task.Do(nil)
}

// go test -v -cover -run=^TestTaskRecovery$
func TestTaskRecovery(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal(r)
		}
	}()

	var task Task = func() { panic("wow") }
	task.Do(func(r any) {
		if r != "wow" {
			t.Fatalf("r %+v is wrong", r)
		}
	})
}
