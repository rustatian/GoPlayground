package main

import "syscall"

func main() {
	_, _, errno := syscall.Syscall6(syscall.Clo, uintptr(0xfee1dead), uintptr(0x28121969), uintptr(0), uintptr(0), 0, 0)
	if errno != 0 {
		panic(errno)
	}
}
