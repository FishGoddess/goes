package lock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

const maxBackoff = 16

type SpinLock uint32

func NewSpinLock() sync.Locker {
	return new(SpinLock)
}

func (sl *SpinLock) Lock() {
	backoff := 1

	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}

		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}
