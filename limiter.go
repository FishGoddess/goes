// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "sync"

const (
	minLimit = 1
	maxLimit = 10000
)

type token struct{}

// Limiter starts some goroutines to do tasks concurrently.
type Limiter struct {
	conf *config

	tokens chan token
	group  sync.WaitGroup
}

// NewLimiter creates a new limiter with limit.
func NewLimiter(limit uint, opts ...Option) *Limiter {
	conf := newConfig().apply(opts...)

	if limit < minLimit {
		limit = minLimit
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return &Limiter{
		conf:   conf,
		tokens: make(chan token, limit),
	}
}

func (l *Limiter) acquireToken() {
	if l.tokens != nil {
		l.tokens <- token{}
	}
}

func (l *Limiter) releaseToken() {
	if l.tokens != nil {
		<-l.tokens
	}
}

// Go starts a goroutine to run task.
func (l *Limiter) Go(task Task) {
	l.acquireToken()
	l.group.Add(1)

	go func() {
		defer func() {
			l.releaseToken()
			l.group.Done()
		}()

		task.Do(l.conf.recovery)
	}()
}

// Wait waits all goroutines to be finished.
func (l *Limiter) Wait() {
	l.group.Wait()
}
