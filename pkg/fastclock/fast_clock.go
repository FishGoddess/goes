// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fastclock

import (
	"sync"
	"sync/atomic"
	"time"
)

const (
	duration   = time.Second
	adjustment = 59
)

// fastClock is a clock for getting current time faster.
// It caches current time in nanos and updates it in fixed duration, so it's not a precise way to get current time.
// In fact, we don't recommend you to use it unless you do need a fast way to get current time even the time is "incorrect".
// According to our benchmarks, it does run faster than time.Now:
//
// In my linux server with 2 cores:
// BenchmarkTimeNow-2               18361782                63.08 ns/op           0 B/op          0 allocs/op
// BenchmarkFastClockNow-2         303623113                 3.92 ns/op           0 B/op          0 allocs/op
// BenchmarkFastClockNowNanos-2    477787526                 2.51 ns/op           0 B/op          0 allocs/op
//
// However, the performance of time.Now is faster enough for 99.9% situations, so we hope you never use it :)
type fastClock struct {
	nanos int64
}

func newClock() *fastClock {
	clock := &fastClock{
		nanos: time.Now().UnixNano(),
	}

	go clock.start()
	return clock
}

func (fc *fastClock) start() {
	for {
		for range adjustment {
			time.Sleep(duration)
			atomic.AddInt64(&fc.nanos, int64(duration))
		}

		time.Sleep(duration)
		atomic.StoreInt64(&fc.nanos, time.Now().UnixNano())
	}
}

func (fc *fastClock) nowNanos() int64 {
	return atomic.LoadInt64(&fc.nanos)
}

var (
	clock     *fastClock
	clockOnce sync.Once
)

// Now returns the current time from fast clock.
func Now() time.Time {
	nanos := NowNanos()
	return time.Unix(0, nanos)
}

// NowNanos returns the current time in nanos from fast clock.
func NowNanos() int64 {
	clockOnce.Do(func() {
		clock = newClock()
	})

	return clock.nowNanos()
}
