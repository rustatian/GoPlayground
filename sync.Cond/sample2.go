package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	s := sync.NewCond(&sync.Mutex{})

	vvv := false
	for !vvv {
		s.L.Lock()
		s.Wait()
	}

	atomic.
		s.L.Unlock()

}
