

package model

import "sync"

// AtomicBool thread safe bool type
type AtomicBool struct {
	bo   bool
	lock sync.Mutex
}

// NewAtomicBool return AtomicBool obj
func NewAtomicBool() *AtomicBool {
	return &AtomicBool{
		bo: false,
	}
}

// Value get AtomicBool value
func (ab *AtomicBool) Value() bool {
	ab.lock.Lock()
	defer ab.lock.Unlock()
	return ab.bo
}

// Set set AtomicBool value
func (ab *AtomicBool) Set(value bool) {
	ab.lock.Lock()
	defer ab.lock.Unlock()
	ab.bo = value
}
