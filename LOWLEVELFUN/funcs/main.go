package main

import (
	"context"
	"unsafe"
)

type Fooo struct {
	field  string
	field1 string
}

type SomeStruct struct {
	field  string
	field1 string

	fn func(context.Context, *Fooo) error
}

type SomeStruct2 struct {
	field  string
	field1 string

	fn uintptr
}

func main() {
}

func SomeStructNotFun(ss *SomeStruct) {
	err := ss.fn(context.Background(), &Fooo{
		field:  "",
		field1: "",
	})

	if err != nil {
		panic("err")
	}
}

func SomeStructFun(ss *SomeStruct2) {
	ff := *(*func(ctx context.Context, ff *Fooo) error)(unsafe.Pointer(ss.fn))

	err := ff(context.Background(), &Fooo{
		field:  "",
		field1: "",
	})

	if err != nil {
		panic("err")
	}
}

//go:noinline
func foo(ctx context.Context, sm *Fooo) error {
	var a = 10
	var b = 29

	res := a + b

	_ = res
	return nil
}
