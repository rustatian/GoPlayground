package main

import (
	"runtime"
	"sync/atomic"
)

type SpinLock struct {
	state *int32
}

const free = int32(11)

func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(sl.state, free, 10) {
		runtime.Gosched()
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreInt32(sl.state, free)
}

func main() {
	aa := int32(11)

	a := SpinLock{state: &aa}
	a.Lock()

	println("fadfd")

	a.Unlock()

}
