package main

import (
	"fmt"
	_ "unsafe"
)

//go:noinline
func Add(x uint64, y uint64) uint64

func main() {
	fmt.Println(Add(1, 2))
}

//go:noinline
func Adding(x uint64, y uint64) uint64 {
	return x + y
}
