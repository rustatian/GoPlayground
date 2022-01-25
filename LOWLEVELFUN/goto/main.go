package main

import (
	"errors"
)

var er = errors.New("foo")

var num uint64 = 0

func main() {
	foo("bbbb")
}

//go:noinline
func foo(a string) {
begin:
	err := foo2()
	if err != nil {
		a = "f"
		goto begin
	}
	return
}

//go:noinline
func foo3(a string) {
	err := foo2()
	if err != nil {
		foo3(a)
	}
	return
}

//go:noinline
func foo2() error {
	if num < 1000 {
		num++
		return er
	}
	return nil
}
