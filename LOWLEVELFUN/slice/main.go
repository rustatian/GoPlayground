package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	b := []byte{1, 2, 4, 5, 6, 6, 6, 6}
	a := conv(b)
	//fmt.Print(a)

	aa := (byte)(uintptr(a[0]))
	aa = 32

	fmt.Print(b)
	fmt.Print(aa)
}

func conv(b []byte) []uint32 {
	val := reflect.ValueOf(b)

	sl := &reflect.SliceHeader{}
	sl.Cap = val.Cap()
	sl.Len = val.Len()
	sl.Data = val.Pointer()

	return *(*[]uint32)(unsafe.Pointer(sl))

}
