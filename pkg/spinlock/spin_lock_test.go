package spinlock

import (
	"sync"
	"testing"
	"time"
)

func benchmarkLock(b *testing.B, lock sync.Locker) {
	b.SetParallelism(1024)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lock.Lock()
			time.Sleep(50 * time.Microsecond)
			lock.Unlock()
		}
	})
}

// go test -v -run=none -bench=^BenchmarkMutex$ -benchmem -benchtime=1s
func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	benchmarkLock(b, &mu)
}

// go test -v -run=none -bench=^BenchmarkSpinLock$ -benchmem -benchtime=1s
func BenchmarkSpinLock(b *testing.B) {
	spin := New()
	benchmarkLock(b, spin)
}
