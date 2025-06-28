// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"
	"time"

	"github.com/FishGoddess/goes/pkg/spinlock"
)

type config struct {
	workerNum        int
	workerQueueSize  int
	nowFunc          func() time.Time
	recoverFunc      func(r any)
	newLockerFunc    func() sync.Locker
	newSchedulerFunc func(workers ...*worker) scheduler
}

func newDefaultConfig(workerNum int) *config {
	return &config{
		workerNum:        workerNum,
		workerQueueSize:  64,
		nowFunc:          nil,
		recoverFunc:      nil,
		newLockerFunc:    nil,
		newSchedulerFunc: nil,
	}
}

func (c *config) now() time.Time {
	if c.nowFunc == nil {
		return time.Now()
	}

	return c.nowFunc()
}

func (c *config) recoverable() bool {
	return c.recoverFunc != nil
}

func (c *config) recover(r any) {
	c.recoverFunc(r)
}

func (c *config) newLocker() sync.Locker {
	if c.newLockerFunc == nil {
		return spinlock.New()
	}

	return c.newLockerFunc()
}

func (c *config) newScheduler(workers ...*worker) scheduler {
	if c.newSchedulerFunc == nil {
		return newRoundRobinScheduler(workers...)
	}

	return c.newSchedulerFunc(workers...)
}
