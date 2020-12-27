package main

import (
	"testing"
)

func Benchmark_Fun(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fun(i)
	}
	b.ResetTimer()
}

func Benchmark_NotFun(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		notfun(i)
	}
	b.ResetTimer()
}
