package main

import (
	"bufio"
	"bytes"
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func Test_superFunc(t *testing.T) {
	a()
}

func a() {
	var buf [8192]byte
	//testing.tRunner(t *T {}, Test_superFunc)
	n := runtime.Stack(buf[:], false)
	sc := bufio.NewScanner(bytes.NewReader(buf[:n]))

	for sc.Scan() {
		var p uintptr
		n, _ := fmt.Sscanf(sc.Text(), "testing.tRunner(%v", &p)
		if n != 1 {
			continue
		}
		//func tRunner(t *T, fn func(t *T))

		var a = (*testing.T)(unsafe.Pointer(p))
		println(a)
	}
}

//goroutine 34 [running]:
//github.com/ValeryPiashchynski/GoPlayground/LOWLEVELFUN/stack/get_stack_vars.a()
///home/valery/Projects/repo/GoPlayground/LOWLEVELFUN/stack/get_stack_vars/main_test.go:24 +0xa2
//github.com/ValeryPiashchynski/GoPlayground/LOWLEVELFUN/stack/get_stack_vars.Test_superFunc(0xc0000fa100)
///home/valery/Projects/repo/GoPlayground/LOWLEVELFUN/stack/get_stack_vars/main_test.go:16 +0x48
//testing.tRunner(0xc0000fa100, 0x5b3f98)
///usr/lib/go/src/testing/testing.go:909 +0x13c
//created by testing.(*T).Run
///usr/lib/go/src/testing/testing.go:960 +0x64f

//goroutine 18 [running]:
//github.com/ValeryPiashchynski/GoPlayground/LOWLEVELFUN/stack/get_stack_vars.a()
///home/valery/Projects/repo/GoPlayground/LOWLEVELFUN/stack/get_stack_vars/main_test.go:24 +0xa2
//github.com/ValeryPiashchynski/GoPlayground/LOWLEVELFUN/stack/get_stack_vars.Test_superFunc(0xc0000d6100)
///home/valery/Projects/repo/GoPlayground/LOWLEVELFUN/stack/get_stack_vars/main_test.go:16 +0x20
//testing.tRunner(0xc0000d6100, 0x5b3f98)
///usr/lib/go/src/testing/testing.go:909 +0x13c
//created by testing.(*T).Run
///usr/lib/go/src/testing/testing.go:960 +0x64f
