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

func (c *config) recover(r any) {
	if c.recoverFunc != nil {
		c.recoverFunc(r)
	}
}
