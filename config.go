// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"

	"github.com/FishGoddess/goes/pkg/spinlock"
)

type config struct {
	size          int
	queueSize     int
	recoverFunc   func(r any)
	newLockerFunc func() sync.Locker
}

func newDefaultConfig(size int) *config {
	return &config{
		size:          size,
		queueSize:     64,
		recoverFunc:   nil,
		newLockerFunc: nil,
	}
}

func (c *config) recover(r any) {
	if c.recoverFunc != nil {
		c.recoverFunc(r)
	}
}

func (c *config) newLocker() sync.Locker {
	if c.newLockerFunc == nil {
		return spinlock.New()
	}

	return c.newLockerFunc()
}
