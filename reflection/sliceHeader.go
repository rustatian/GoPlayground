package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	sliceOne := []int{0, 1, 2, 4, 5}
	sliceTwo := []int{6, 7, 8, 9, 10}

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&sliceOne))
	sh2 := (*reflect.SliceHeader)(unsafe.Pointer(&sliceTwo))
	sh1.Data = sh2.Data

	for i := 0; i < len(sliceOne); i++ {
		fmt.Println(sliceOne[i])
		fmt.Println(sliceTwo[i])
	}
}
