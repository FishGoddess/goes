// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import "sync"

const (
	minLimit = 1
	maxLimit = 10000
)

type token struct{}

// Limiter limits the simultaneous number of goroutines.
type Limiter struct {
	tokens chan token
	wg     sync.WaitGroup
}

// NewLimiter creates a new limiter with limit.
func NewLimiter(limit int) *Limiter {
	if limit < minLimit {
		limit = minLimit
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return &Limiter{
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

// Go starts a goroutine with token.
func (l *Limiter) Go(f func()) {
	l.acquireToken()
	l.wg.Add(1)

	go func() {
		defer func() {
			l.releaseToken()
			l.wg.Done()
		}()

		f()
	}()
}

// Wait waits all goroutines to be finished.
func (l *Limiter) Wait() {
	l.wg.Wait()
}
