package spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

const maxBackoff = 16

type Lock uint32

// New creates a new spin lock.
func New() sync.Locker {
	return new(Lock)
}

// Lock locks with the spin lock.
func (l *Lock) Lock() {
	backoff := 1

	for !atomic.CompareAndSwapUint32((*uint32)(l), 0, 1) {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}

		if backoff < maxBackoff {
			backoff <<= 1
		}
	}
}

// Unlock unlocks with the spin lock.
func (l *Lock) Unlock() {
	atomic.StoreUint32((*uint32)(l), 0)
}
