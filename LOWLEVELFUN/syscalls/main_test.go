package main

import (
	"testing"
	"unsafe"
)

var lens int = 1000000000
var size *int

func Benchmark_Funslice(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		a := make_slice_non_fun(lens, unsafe.Sizeof(size))
		_ = a
	}
}

func Benchmark_NotFunslice(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		a := make_slice_non_fun(lens, 0)
		_ = a
	}
}
