package main

import (
	"crypto/rand"
	"testing"
)

func Benchmark_Val1(b *testing.B) {
	b1 := make([]byte, 1024)
	b2 := make([]byte, 1024)

	_, err := rand.Read(b1)
	if err != nil {
		b.Fatal(err)
	}

	_, err = rand.Read(b2)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	val := Payload{
		Context: b1,
		Body:    b2,
	}

	for i := 0; i < b.N; i++ {
		l := fooVal1(val)
		_ = l
	}
}

func Benchmark_Ptr1(b *testing.B) {
	b1 := make([]byte, 1024)
	b2 := make([]byte, 1024)

	_, err := rand.Read(b1)
	if err != nil {
		b.Fatal(err)
	}

	_, err = rand.Read(b2)
	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	pld := &Payload{
		Context: b1,
		Body:    b2,
	}

	for i := 0; i < b.N; i++ {
		l := fooPtr1(pld)
		_ = l
	}
}
