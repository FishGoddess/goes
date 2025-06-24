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

// WithQueueSize sets the queue size of worker.
func WithQueueSize(size int) Option {
	return func(conf *config) {
		conf.queueSize = size
	}
}

// WithRecoverFunc sets the recover function.
func WithRecoverFunc(f func(r any)) Option {
	return func(conf *config) {
		conf.recoverFunc = f
	}
}

// WithNewLockerFunc sets the new locker function.
func WithNewLockerFunc(f func() sync.Locker) Option {
	return func(conf *config) {
		conf.newLockerFunc = f
	}
}
