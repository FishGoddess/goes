// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

type config struct {
	size        int
	queueSize   int
	recoverFunc func(r any)
}

func newDefaultConfig(size int) *config {
	return &config{
		size:        size,
		queueSize:   64,
		recoverFunc: nil,
	}
}

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

// WithRecoverFunc sets the recover function of worker.
func WithRecoverFunc(f func(r any)) Option {
	return func(conf *config) {
		conf.recoverFunc = f
	}
}
