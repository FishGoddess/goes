// Copyright 2023 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goes

import (
	"testing"
	"time"
)

// go test -v -cover -run=^TestLimiter$
func TestLimiter(t *testing.T) {
	limiter := NewLimiter(4)

	for i := 0; i < 100; i++ {
		limiter.Go(func() {
			t.Log(time.Now())
			time.Sleep(10 * time.Millisecond)
		})
	}

	limiter.Wait()
}
