package main

import (
	"fmt"
	_ "unsafe"
)



func Add(x uint64, y uint64) uint64

func main() {
	//	TEXT("Add", NOSPLIT, "func(x, y uint64) uint64")
	//	Doc("Add adds x and y.")
	//	x := Load(Param("x"), GP64())
	//	y := Load(Param("y"), GP64())
	//	ADDQ(x, y)
	//	Store(y, ReturnIndex(0))
	//	RET()
	//	Generate()

	fmt.Println(Add(1, 2))
}

//go:noinline
func Adding(x uint64, y uint64) uint64 {
	return x + y
}
