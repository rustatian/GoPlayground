package main

import "testing"

func BenchmarkAnd(b *testing.B) {
	fi := &FooImpl{}

	for i := 0; i < b.N; i++ {
		res := fi.State() & 1
		_ = res
		if (fi.State() & 1) == 0 {
			continue
		}
	}
}

func BenchmarkEq(b *testing.B) {
	fi := &FooImpl{}

	for i := 0; i < b.N; i++ {
		if fi.State() == 1 {
			continue
		}
	}
}
