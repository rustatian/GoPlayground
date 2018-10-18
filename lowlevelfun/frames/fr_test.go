package ttt

import (
	"testing"
)

func Benchmark_Pointer(b *testing.B) {
	c := callers()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = frameP(c, 2)
	}
}

func Benchmark_Value(b *testing.B) {
	c := callers()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_ = frameV(c, 2)
	}
}
