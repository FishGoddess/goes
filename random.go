// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"math/rand"
	"time"
)

type randomScheduler struct {
	workers []*worker
	random  *rand.Rand
}

func newRandomScheduler(workers ...*worker) *randomScheduler {
	scheduler := &randomScheduler{
		workers: workers,
		random:  rand.New(rand.NewSource(time.Now().Unix())),
	}

	return scheduler
}

// Set sets the workers to scheduler.
func (rs *randomScheduler) Set(workers []*worker) {
	rs.workers = workers
}

// Get gets a worker from scheduler.
func (rs *randomScheduler) Get() *worker {
	if len(rs.workers) <= 0 {
		return nil
	}

	index := rs.random.Intn(len(rs.workers))
	worker := rs.workers[index]
	return worker
}
