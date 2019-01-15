package main

import "testing"


func BenchmarkAddNonFun(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i ++ {
		x := Adding(uint64(i),2)
		_ = x
	}
}

func BenchmarkAddFun(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i ++ {
		x := Add(uint64(i),2)
		_ = x
	}
}
