// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestLimiter$
func TestLimiter(t *testing.T) {
	limit := 16
	limiter := NewLimiter(limit)

	var countMap = make(map[int64]int, 16)
	var lock sync.Mutex

	totalCount := 100
	for i := 0; i < totalCount; i++ {
		limiter.Go(func() {
			now := time.Now().UnixMilli()

			lock.Lock()
			countMap[now] = countMap[now] + 1
			lock.Unlock()

			time.Sleep(100 * time.Millisecond)
		})
	}

	limiter.Wait()

	gotTotalCount := 0
	for now, count := range countMap {
		gotTotalCount = gotTotalCount + count

		if count != limit {
			t.Logf("now %d: count %d != limit %d", now, count, limit)
		}
	}

	if gotTotalCount != totalCount {
		t.Fatalf("gotTotalCount %d != totalCount %d", gotTotalCount, totalCount)
	}
}
