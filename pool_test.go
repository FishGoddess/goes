// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestPool$
func TestPool(t *testing.T) {
	size := 16
	queueSize := 1024
	pool := NewPool(size, WithQueueSize(queueSize))

	var countMap = make(map[int64]int, 16)
	var lock sync.Mutex

	totalCount := 10 * size
	for i := 0; i < totalCount; i++ {
		pool.Submit(func() {
			now := time.Now().UnixMilli() / 10

			lock.Lock()
			countMap[now] = countMap[now] + 1
			lock.Unlock()

			time.Sleep(10 * time.Millisecond)
		})
	}

	pool.Close()
	pool.Wait()

	gotTotalCount := 0
	for now, count := range countMap {
		gotTotalCount = gotTotalCount + count

		if count != size {
			t.Fatalf("now %d: count %d != size %d", now, count, size)
		}
	}

	if gotTotalCount != totalCount {
		t.Fatalf("gotTotalCount %d != totalCount %d", gotTotalCount, totalCount)
	}
}
