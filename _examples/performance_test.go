// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/FishGoddess/goes"
	//"github.com/panjf2000/ants/v2"
)

const (
	limit    = 256
	size     = limit
	timeLoop = 100_0000
)

func bench(num *uint32) {
	atomic.AddUint32(num, 1)
}

// go test -v -run=none -bench=^BenchmarkLimiter$ -benchmem -benchtime=1s
func BenchmarkLimiter(b *testing.B) {
	limiter := goes.NewLimiter(limit)

	num := uint32(0)
	f := func() {
		bench(&num)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Go(f)
		}
	})

	limiter.Wait()
	b.Logf("num is %d", num)
}

// go test -v -run=none -bench=^BenchmarkLimiterTime$ -benchmem -benchtime=1s
func BenchmarkLimiterTime(b *testing.B) {
	limiter := goes.NewLimiter(limit)

	num := uint32(0)
	f := func() {
		bench(&num)
	}

	beginTime := time.Now()
	for range timeLoop {
		limiter.Go(f)
	}

	limiter.Wait()

	cost := time.Since(beginTime)
	b.Logf("num is %d, cost is %s", num, cost)
}

// go test -v -run=none -bench=^BenchmarkPool$ -benchmem -benchtime=1s
func BenchmarkPool(b *testing.B) {
	workerLimit := 1024
	pool := goes.NewPool(size, workerLimit)

	num := uint32(0)
	task := func() {
		bench(&num)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pool.Submit(task)
		}
	})

	pool.Close()
	b.Logf("num is %d", num)
}

// go test -v -run=none -bench=^BenchmarkPoolTime$ -benchmem -benchtime=1s
func BenchmarkPoolTime(b *testing.B) {
	workerLimit := 1024
	pool := goes.NewPool(size, workerLimit)

	num := uint32(0)
	task := func() {
		bench(&num)
	}

	beginTime := time.Now()
	for range timeLoop {
		pool.Submit(task)
	}

	pool.Close()

	cost := time.Since(beginTime)
	b.Logf("num is %d, cost is %s", num, cost)
}

// // go test -v -run=none -bench=^BenchmarkAntsPool$ -benchmem -benchtime=1s
// func BenchmarkAntsPool(b *testing.B) {
// 	pool, _ := ants.NewPool(size)
//
// 	num := uint32(0)
// 	task := func() {
// 		bench(&num)
// 	}
//
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			pool.Submit(task)
// 		}
// 	})
//
// 	pool.Release()
// 	b.Logf("num is %d", num)
// }
//
// // go test -v -run=none -bench=^BenchmarkAntsPoolTime$ -benchmem -benchtime=1s
// func BenchmarkAntsPoolTime(b *testing.B) {
// 	pool, _ := ants.NewPool(size)
//
// 	num := uint32(0)
// 	task := func() {
// 		bench(&num)
// 	}
//
// 	beginTime := time.Now()
// 	for range timeLoop {
// 		pool.Submit(task)
// 	}
//
// 	pool.Release()
//
// 	cost := time.Since(beginTime)
// 	b.Logf("num is %d, cost is %s", num, cost)
// }
