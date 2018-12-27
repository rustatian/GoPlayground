package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	var example *int
	slice := make_slice(1e9, unsafe.Sizeof(example))
	a := *(*[]*int)(unsafe.Pointer(&slice))

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	runtime.KeepAlive(a)
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func make_slice(len int, eltsize uintptr) SliceHeader {
	fd := -1
	data, _, errno := syscall.Syscall6(
		syscall.SYS_MMAP,
		0,
		uintptr(len)*eltsize,
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_ANON|syscall.MAP_PRIVATE,
		uintptr(fd),
		0,
	)

	if errno != 0 {
		panic(errno)
	}
	return SliceHeader{
		Data: data,
		Len:  len,
		Cap:  len,
	}
}
