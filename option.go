// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "sync"

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

// WithSyncMutex sets the new locker function returns sync.Mutex.
func WithSyncMutex() Option {
	newLockerFunc := func() sync.Locker {
		return new(sync.Mutex)
	}

	return func(conf *config) {
		conf.newLockerFunc = newLockerFunc
	}
}
