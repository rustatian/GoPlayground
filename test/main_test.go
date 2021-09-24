package main

import (
	"runtime"
	"testing"
)

func BenchmarkSleep(b *testing.B) {
	runtime.GOMAXPROCS(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sleep()
	}
}

func BenchmarkBlock(b *testing.B) {
	runtime.GOMAXPROCS(1)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		block()
	}
}
