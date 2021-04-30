package main

import (
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	s := &Shared{
		data:      0,
		lockValue: 0,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.test_mutexes()
	}
}

func BenchmarkAtomic(b *testing.B) {
	s := &Shared{
		data:      0,
		lockValue: 0,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.test_atomic()
	}
}
