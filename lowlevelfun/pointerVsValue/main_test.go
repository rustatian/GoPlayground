package main

import "testing"

type TL [2]uintptr

func (TL) Method1() {

}

func (TL) Method2() {

}

var tl TL

var e interface{}

func BenchmarkA_FooP(b *testing.B) {
	aa := aa{}
	a := &foos{
		a: "yuv has", b: "a small", c: "dick",
		d: struct {
			a string
			b string
			c string
		}{a: "really", b: "really", c: "small"},
		foo: aa,
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		FooP(a)
	}
}

func BenchmarkA_FooV(b *testing.B) {
	aa := aa{}
	a := foos{
		a: "yuv has", b: "a small", c: "dick",
		d: struct {
			a string
			b string
			c string
		}{a: "really", b: "really", c: "small"},
		foo: aa,
	}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		FooV(a)
	}
}
