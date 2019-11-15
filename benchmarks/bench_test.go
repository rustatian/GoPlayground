package main

import "testing"

func Benchmark_One(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ii := One()
		_ = ii
	}
}

func Benchmark_Two(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Two()
	}
}
