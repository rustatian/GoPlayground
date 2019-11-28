package main

import (
	"fmt"
	"unsafe"
)

type foo interface {
	FooErr()
	Foo()
}

type foos struct {
	a int
	b int
	c string
	d struct {
		a int
		b int
		c string
	}

	foo
}

func main() {
	b := &aa{}

	a := &foos{
		//a: 1, b: 2, c: "fsf",
		//d: struct {
		//	a int
		//	b int
		//	c string
		//}{a: 1, b: 2, c: "33"},

		foo: b,
	}

	aa := foos{
		//a: 1, b: 2, c: "fsf",
		//d: struct {
		//	a int
		//	b int
		//	c string
		//}{a: 1, b: 2, c: "33"},

		foo: b,
	}

	fmt.Println(unsafe.Sizeof(aa))

	fmt.Println(unsafe.Sizeof(a))

}

type aa struct {
}

func (aa) FooErr() {
	panic("implement me")
}

func (aa) Foo() {
	panic("implement me")
}
