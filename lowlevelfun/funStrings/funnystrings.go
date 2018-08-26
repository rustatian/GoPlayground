package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print(FunConvert("fdsf"))
}

func FunConvert(intr interface{}) []byte {
	return *(*[]byte)(unsafe.Pointer(&intr))
}

func NonFunConvert(intr interface{}) string {
	return intr.(string)
}
