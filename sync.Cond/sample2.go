package main

import (
	"sync"
)

func main() {
	s := sync.NewCond(&sync.Mutex{})

	vvv := false
	for !vvv {
		s.L.Lock()
		s.Wait()
	}
}
