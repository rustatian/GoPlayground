package main

import (
	"sync"
	"testing"
)

var a string

func BenchmarkRWMutex(t *testing.B) {
	mu := &sync.RWMutex{}

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		mu.Lock()
		a = "foo"
		_ = a
		mu.Unlock()
	}
}

func BenchmarkMutex(t *testing.B) {
	mu := &sync.Mutex{}

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		mu.Lock()
		a = "foo"
		_ = a
		mu.Unlock()
	}
}

func BenchmarkChannel(t *testing.B) {
	ch := make(chan struct{}, 1)

	t.ResetTimer()
	t.ReportAllocs()
	for i := 0; i < t.N; i++ {
		ch <- struct{}{}
		a = "foo"
		_ = a
		<-ch
	}
}
