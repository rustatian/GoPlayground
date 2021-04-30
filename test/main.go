package main

import (
	"sync"
	"sync/atomic"
)

func main() {

}

type Shared struct {
	sync.RWMutex
	data      uint64
	lockValue uint64
}

//go:noinline
func (s *Shared) test_mutexes() {
	s.Lock()
	s.data = 132
	s.Unlock()
}

//go:noinline
func (s *Shared) test_atomic() {
	if atomic.CompareAndSwapUint64(&s.lockValue, 0, 0) {
		s.data = 132
	}
}
