// Copyright 2025 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"context"
	"sync"
	"testing"
	"time"
)

// go test -v -cover -run=^TestExecutor$
func TestExecutor(t *testing.T) {
	workers := 16
	executor := NewExecutor(uint(workers))
	defer executor.Close()

	var countMap = make(map[int64]int, 16)
	var lock sync.Mutex

	ctx := context.Background()
	totalCount := 10 * workers
	for i := 0; i < totalCount; i++ {
		executor.Submit(ctx, func() {
			now := time.Now().UnixMilli() / 10

			lock.Lock()
			countMap[now] = countMap[now] + 1
			lock.Unlock()

			time.Sleep(10 * time.Millisecond)
		})
	}

	executor.Close()

	gotTotalCount := 0
	for now, count := range countMap {
		gotTotalCount = gotTotalCount + count

		if count != workers {
			t.Fatalf("now %d: count %d != workers %d", now, count, workers)
		}
	}

	if gotTotalCount != totalCount {
		t.Fatalf("gotTotalCount %d != totalCount %d", gotTotalCount, totalCount)
	}
}
