// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "testing"

// go test -v -cover -run=^TestNewDefaultConfig$
func TestNewDefaultConfig(t *testing.T) {
	size := 16
	conf := newDefaultConfig(size)

	if conf.size != size {
		t.Fatalf("conf.size %d != size %d", conf.size, size)
	}
}

// go test -v -cover -run=^TestConfigRecover$
func TestConfigRecover(t *testing.T) {
	size := 16
	conf := newDefaultConfig(size)
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
