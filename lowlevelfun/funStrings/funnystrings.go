package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Print()
	up := FromInterfaceToBytes([]byte{0x11, 0x23, 0x34})
	fmt.Print(up)

	//type MyString string
	//ms := []MyString{"C", "C++", "Go"}
	//fmt.Printf("%s\n", ms)  // [C C++ Go]
	//// ss := ([]string)(ms) // compiling error
	//ss := *(*[]string)(unsafe.Pointer(&ms))
	//ss[1] = "Rust"
	//fmt.Printf("%s\n", ms) // [C Rust Go]
	//// ms = []MyString(ss) // compiling error
	//ms = *(*[]MyString)(unsafe.Pointer(&ss))
}

func FromInterfaceToBytes(intr interface{}) []byte {
	p := *(*interface{})(unsafe.Pointer(&intr))
	b := p.([]byte)
	return b
}

func FunConvert(intr interface{}) unsafe.Pointer {
	return unsafe.Pointer(&intr)
}

func NonFunConvert(intr interface{}) string {
	return intr.(string)
}
