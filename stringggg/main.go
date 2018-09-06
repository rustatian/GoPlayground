package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

func main() {
	s := "hello"
	header := (*StringHeader)(unsafe.Pointer(&s))
	header.Len = 100000

	fmt.Print(*(*string)(unsafe.Pointer(header)))

}
