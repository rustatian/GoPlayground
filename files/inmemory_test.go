package main

import "testing"

func Benchmark_1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandStringBytesMaskImprSrc(i)
	}

}

func Benchmark_2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandBytes(i)
	}
}
