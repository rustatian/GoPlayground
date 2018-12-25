package main

import "testing"

func Benchmark_map(b *testing.B) {
	b.ReportAllocs()
	for i :=0; i < b.N; i ++ {
		x := mmap()
		_ = x
	}
}

func Benchmark_new(b *testing.B) {
	b.ReportAllocs()
	for i :=0; i < b.N; i ++ {
		x := nnew()
		_ = x
	}
}
