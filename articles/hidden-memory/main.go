package main

import (
	"fmt"
	"runtime"
	"time"
)

var sliceLen int = 10000000

type foo struct {
	pSlice []*int
	//vSlice []int
}

func main() {
	f := &foo{
		pSlice: make([]*int, sliceLen),
		//vSlice: make([]int, sliceLen),
	}

	runtime.KeepAlive(f)

	for i := 0; i <= 10; i++ {
		start := time.Now()
		runtime.GC()

		fmt.Printf("GC time %s\n", time.Since(start))
	}

}
