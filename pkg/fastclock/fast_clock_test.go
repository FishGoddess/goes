// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fastclock

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

// go test -v -cover -run=^TestNow$
func TestNow(t *testing.T) {
	for i := 0; i < 10; i++ {
		got := Now()
		gap := time.Since(got)
		t.Logf("got: %v, gap: %v", got, gap)

		if math.Abs(float64(gap.Nanoseconds())) > float64(duration)*1.2 {
			t.Errorf("now %v is wrong", got)
		}

		time.Sleep(time.Duration(rand.Int63n(int64(duration))))
	}
}

// go test -v -cover -run=^TestNowNanos$
func TestNowNanos(t *testing.T) {
	for i := 0; i < 10; i++ {
		gotNanos := NowNanos()
		got := time.Unix(0, gotNanos)
		gap := time.Since(got)
		t.Logf("got: %v, gap: %v", got, gap)

		if math.Abs(float64(gap.Nanoseconds())) > float64(duration)*1.2 {
			t.Errorf("now %v is wrong", got)
		}

		time.Sleep(time.Duration(rand.Int63n(int64(duration))))
	}
}

// go test -v -run=none -bench=^BenchmarkTimeNow$ -benchmem -benchtime=1s
func BenchmarkTimeNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

// go test -v -run=none -bench=^BenchmarkFastClockNow$ -benchmem -benchtime=1s
func BenchmarkFastClockNow(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Now()
	}
}

// go test -v -run=none -bench=^BenchmarkFastClockNowNanos$ -benchmem -benchtime=1s
func BenchmarkFastClockNowNanos(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NowNanos()
	}
}
