package main

import (
	"fmt"
	"sync/atomic"
)

type foo struct {
	C bool
	A string
	B int
}

func main() {
	//sync.Mutex{}
	aa := atomic.Value{}
	aa.Store("fsfd")
	vv := aa.Load()
	//aa := unsafe.Sizeof(nil)
	fmt.Print(vv)
}
