package main

import (
	_ "unsafe"
)

func main() {
	id := runtime_procPin()
	println(id)
	runtime_procUnpin()
}

//go:linkname runtime_procPin runtime.procPin
func runtime_procPin() int

//go:linkname runtime_procUnpin runtime.procUnpin
func runtime_procUnpin() int
