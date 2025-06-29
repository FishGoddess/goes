// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"
	"time"

	"github.com/FishGoddess/goes/pkg/spinlock"
)

// Option is for setting config.
type Option func(conf *config)

func (o Option) applyTo(conf *config) {
	o(conf)
}

// WithWorkerQueueSize sets the queue size of worker.
func WithWorkerQueueSize(size int) Option {
	return func(conf *config) {
		conf.workerQueueSize = size
	}
}

// WithPurgeActive sets the purge interval of executor and the lifetime of worker.
func WithPurgeActive(purgeInterval time.Duration, workerLifetime time.Duration) Option {
	return func(conf *config) {
		conf.purgeInterval = purgeInterval
		conf.workerLifetime = workerLifetime
	}
}

// WithNowFunc sets the now function.
func WithNowFunc(nowFunc func() time.Time) Option {
	return func(conf *config) {
		conf.nowFunc = nowFunc
	}
}

// WithRecoverFunc sets the recover function.
func WithRecoverFunc(recoverFunc func(r any)) Option {
	return func(conf *config) {
		conf.recoverFunc = recoverFunc
	}
}

// WithNewLockerFunc sets the new locker function.
func WithNewLockerFunc(newLockerFunc func() sync.Locker) Option {
	return func(conf *config) {
		conf.newLockerFunc = newLockerFunc
	}
}

// WithSpinLock sets the new locker function returns spin lock.
func WithSpinLock() Option {
	newLockerFunc := func() sync.Locker {
		return spinlock.New()
	}

	return func(conf *config) {
		conf.newLockerFunc = newLockerFunc
	}
}

// WithSyncMutex sets the new locker function returns sync mutex.
func WithSyncMutex() Option {
	newLockerFunc := func() sync.Locker {
		return new(sync.Mutex)
	}

	return func(conf *config) {
		conf.newLockerFunc = newLockerFunc
	}
}

// WithRoundRobinScheduler sets the new scheduler function using round robin strategy.
func WithRoundRobinScheduler() Option {
	newSchedulerFunc := func(workers ...*worker) scheduler {
		return newRoundRobinScheduler(workers...)
	}

	return func(conf *config) {
		conf.newSchedulerFunc = newSchedulerFunc
	}
}

// WithRandomScheduler sets the new scheduler function using random strategy.
func WithRandomScheduler() Option {
	newSchedulerFunc := func(workers ...*worker) scheduler {
		return newRandomScheduler(workers...)
	}

	return func(conf *config) {
		conf.newSchedulerFunc = newSchedulerFunc
	}
}
