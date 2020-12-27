package main

import (
	"unsafe"
)

func main() {
	s := "abcdefgggggg"

	sss := *(*StringHeader)(unsafe.Pointer(&s))

	println(sss.Len)
	println(*(*string)(unsafe.Pointer(&sss.Data)))

	println("----------------------------------------------------------------------")

	sss.Len = 1000
	//
	println(*(*string)(unsafe.Pointer(&sss.Data)))

}

type StringHeader struct {
	Data uintptr
	Len  int
}
