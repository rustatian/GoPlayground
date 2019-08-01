package main

import (
	"sort"
	"testing"
)

func Benchmark_naive_search(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = naive_search(targetSlice,searchStrin)

	}
}

func Benchmark_sorted_search(b *testing.B) {
	sort.Strings(targetSlice)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = sorted_search(targetSlice,searchStrin)

	}
}
