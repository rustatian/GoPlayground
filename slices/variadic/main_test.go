package main

import (
	"math/rand"
	"testing"
)

func BenchmarkA(b *testing.B) {
	data := make([]*int, 1000)

	for i := 0; i < 1000; i++ {
		data[i] = toPtr(rand.Int())
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := A(data...)
		_ = sum
	}
}

func BenchmarkB(b *testing.B) {
	data := make([]*int, 1000)

	for i := 0; i < 1000; i++ {
		data[i] = toPtr(rand.Int())
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := B(data)
		_ = sum
	}
}

func toPtr(i int) *int {
	return &i
}
