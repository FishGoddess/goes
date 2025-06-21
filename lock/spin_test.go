package lock

import (
	"sync"
	"testing"
	"time"
)

func doSomething() {
	time.Sleep(50 * time.Microsecond)
}

func BenchmarkMutex(b *testing.B) {
	mu := sync.Mutex{}

	b.ReportAllocs()
	b.SetParallelism(1024)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			doSomething()
			mu.Unlock()
		}
	})
}

func BenchmarkBackoffSpinLock(b *testing.B) {
	spin := NewSpinLock()

	b.ReportAllocs()
	b.SetParallelism(1024)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			spin.Lock()
			doSomething()
			spin.Unlock()
		}
	})
}
