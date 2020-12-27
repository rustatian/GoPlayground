package main

import "testing"

func Benchmark_optimized(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		a := returnMaxIntOptimized()
		_ = a
	}
}

func Benchmark__not_optimized(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		a := returnMaxIntNotOptimized()
		_ = a
	}
}
