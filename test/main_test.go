package main

import (
	"context"
	"testing"
)

func BenchmarkMy(b *testing.B) {
	bgr := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.WithValue(bgr, PsrContextKey, "foo")
		_ = ctx.Value("foo")
	}
}

func BenchmarkNr(b *testing.B) {
	bgr := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.WithValue(bgr, GinTransactionContextKey, "foo")
		_ = ctx.Value("foo")
	}
}
