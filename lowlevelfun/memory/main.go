package main

import (
	"runtime"
)

func main() {
	a := make([]*int, 100)
	runtime.GC()
	runtime.KeepAlive(a)
}
