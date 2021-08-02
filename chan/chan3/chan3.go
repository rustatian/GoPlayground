package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ch Chan
	go func() { ch.Get() <- struct{}{} }()
	<-ch.Get()

	fmt.Println("it works?")
}

// this has to be kept in line with runtime.hchan
//
// it turns out that most of the code in makechan is a no-op for
// chan struct{} with no capacity, so we can actually get away
// with using the zero value.
//
// one complication is the goexperiment.staticlockranking build
// tag. if it is specified, the mutex type changes in size, and
// unfortunately the futex value moves. the comment in the runtime
// though says
//
//	Initialization is helpful for static lock ranking, but not required.
//
// so the struct has been padded with the worst case size including
// the 2 int fields, and the mutex will just use some of that memory.
// it's no big deal if our Chan is bigger than the standard runtime.hchan
// because it turns out that in some cases the buffer is allocated
// with the runtime.hchan value.

type Chan struct {
	_ uint
	_ uint
	_ unsafe.Pointer
	_ uint16
	_ uint32
	_ unsafe.Pointer // *_type
	_ uint
	_ uint
	_ struct {
		_ unsafe.Pointer // *sudog
		_ unsafe.Pointer // *sudog
	} // waitq
	_ struct {
		_ unsafe.Pointer // *sudog
		_ unsafe.Pointer // *sudog
	} // waitq

	_ struct {
		_ [2]int // worst case lock ranking
		_ uintptr
	} // mutex
}

func (c *Chan) Get() chan struct{} {
	return *(*chan struct{})(unsafe.Pointer(&c))
}
