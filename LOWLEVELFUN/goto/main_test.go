package main

import (
	"testing"
)

func BenchmarkRecursion(b *testing.B) {
	num = 0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		foo3("bar")
	}
}

func BenchmarkGoTo(b *testing.B) {
	num = 0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		foo("bar")
	}
}
