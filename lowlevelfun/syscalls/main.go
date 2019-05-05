package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	//var example *string
	//slice := make_slice_non_fun(1e5, unsafe.Sizeof(example))
	slice := make([]*string, 1e5, 1e5)
	//a := *(*[]*int)(unsafe.Pointer(&slice))

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s\n", time.Since(start))
	}

	runtime.KeepAlive(slice)
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

//http://man7.org/linux/man-pages/man2/mmap.2.html
//void *mmap(void *addr, size_t length, int prot, int flags, int fd, off_t offset);

func make_slice(len int, eltsize uintptr) []*string {
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

	slice := SliceHeader{
		Data: data,
		Len:  len,
		Cap:  len,
	}

	return *(*[]*string)(unsafe.Pointer(&slice))
}

func make_slice_non_fun(len int, eltsize uintptr) []*string {
	slice := make([]*string, len, len)
	return slice
}
