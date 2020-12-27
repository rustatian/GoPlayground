package main

import (
	"reflect"
	"unsafe"
)

func main() {
	ss := "waaaatttttttttttttttttttttt"
	//change_by_pointer(&ss)
	magic_change(ss)
	println(ss)
}

//go:noinline
func change_by_pointer(s *string) {
	*s = "whatever"
}

func magic_change(s string) {
	//sss := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	//println(*(*string)(unsafe.Pointer(&sss.Data)))
	//
	//a := *(*string)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data))

	//newStd := "wwwhhhhhaaatttt"
	//(*string)(unsafe.Pointer(&(*reflect.StringHeader)(unsafe.Pointer(&s)).Data)) = &newStd
	*(*int)(unsafe.Pointer(&(*reflect.StringHeader)(unsafe.Pointer(&s)).Len)) = len("whatever")

	newStr := "whatever"
	data := []byte{'w'}
	sss := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	sss.Len = len(newStr)
	sss.Data = uintptr(unsafe.Pointer(&data))

	//println(*(*string)(unsafe.Pointer(&sss.Data)))
}

func foo(a []bool) []bool {
	a = nil
	return a
}
