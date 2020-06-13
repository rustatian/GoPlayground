package main

import (
	"reflect"
	"unsafe"
)

type S1 struct {

}

func (s *S1) Init() error {
	println("hello")
	return nil
}

func funcAddr(fn interface{}) uintptr {
	// emptyInterface is the header for an interface{} value.
	type emptyInterface struct {
		typ   uintptr
		value *uintptr
	}
	e := (*emptyInterface)(unsafe.Pointer(&fn))
	return *e.value
}

func main() {
	//s := *S1{}
	v := reflect.Value{}.Set

	addr := funcAddr((*S1).Init)
	v.Pointer = addr
	println(addr)
}
