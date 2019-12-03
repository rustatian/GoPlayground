package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Print()
	up := FromInterfaceToBytes("some text for parse")
	fmt.Print(up)

	runtime.Gosched()
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
