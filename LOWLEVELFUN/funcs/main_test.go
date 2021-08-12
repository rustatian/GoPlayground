package main

import (
	"testing"
	"unsafe"
)

func BenchmarkSomeStructNotFun(b *testing.B) {
	s := &SomeStruct{
		field:  "aaaaaaaaaaaaaaaa",
		field1: "bbbbbbbbbbbbbbbb",
		fn:     foo,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		SomeStructNotFun(s)
	}
}

func BenchmarkSomeStructFun(b *testing.B) {
	fn := foo
	s := &SomeStruct2{
		field:  "aaaaaaaaaaaaaaaa",
		field1: "bbbbbbbbbbbbbbbb",
		fn:     uintptr(unsafe.Pointer(&fn)),
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		SomeStructFun(s)
	}
}
