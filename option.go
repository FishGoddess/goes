// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"time"
)

type config struct {
	workerNum       int
	workerQueueSize int
	now             func() time.Time
	handlePanic     func(r any)
}

func newConfig(workerNum int) *config {
	return &config{
		workerNum:       workerNum,
		workerQueueSize: 256,
		now:             time.Now,
		handlePanic:     nil,
	}
}

type Option func(conf *config)

// WithWorkerQueueSize sets the queue size of worker.
func WithWorkerQueueSize(size int) Option {
	return func(conf *config) {
		conf.workerQueueSize = size
	}
}

// WithNow sets the now function.
func WithNow(now func() time.Time) Option {
	return func(conf *config) {
		conf.now = now
	}
}

// WithHandlePanic sets the handle panic function.
func WithHandlePanic(handlePanic func(r any)) Option {
	return func(conf *config) {
		conf.handlePanic = handlePanic
	}
}
