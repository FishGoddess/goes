// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type config struct {
	queueSize uint
	recovery  func(r any)
}

func newConfig() *config {
	return &config{
		queueSize: 64,
		recovery:  nil,
	}
}

func (c *config) apply(opts ...Option) *config {
	for _, opt := range opts {
		opt(c)
	}

	return c
}

type Option func(conf *config)

// WithQueueSize sets the queue size of worker.
func WithQueueSize(queueSize uint) Option {
	return func(conf *config) {
		conf.queueSize = queueSize
	}
}

// WithRecovery sets the recovery function.
func WithRecovery(recovery func(r any)) Option {
	return func(conf *config) {
		conf.recovery = recovery
	}
}
