package main

import (
	"crypto/rand"
	"testing"
)

func BenchmarkA_FooP(b *testing.B) {
	b1 := make([]byte, 351024)
	b2 := make([]byte, 351024)

	_, err := rand.Read(b1)
	if err != nil {
		b.Fatal(err)
	}

	_, err = rand.Read(b2)
	if err != nil {
		b.Fatal(err)
	}
	a := &foos{
		a:  "yuv has",
		b:  "a small",
		c:  "dick",
		f:  "1 mmmmmmiiiiiilllimeter",
		g:  "or might be 2",
		h:  "you know yuv",
		hh: "he is from the project",
		aa: "IN",
		bb: "TU",
		cc: "RN",
		ee: string(b1),
		dd: string(b2),
		d: struct {
			a string
			b string
			c string
		}{a: "really", b: "really", c: "small"},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		ret := FooP(a)
		_ = ret
	}
}

func BenchmarkA_FooV(b *testing.B) {
	b1 := make([]byte, 351024)
	b2 := make([]byte, 351024)

	_, err := rand.Read(b1)
	if err != nil {
		b.Fatal(err)
	}

	_, err = rand.Read(b2)
	if err != nil {
		b.Fatal(err)
	}

	a := foos{
		a:  "yuv has",
		b:  "a small",
		c:  "dick",
		f:  "1 mmmmmmiiiiiilllimeter",
		g:  "or might be 2",
		h:  "you know yuv",
		hh: "he is from the project",
		aa: "IN",
		bb: "TU",
		cc: "RN",
		ee: string(b1),
		dd: string(b2),
		d: struct {
			a string
			b string
			c string
		}{a: "really", b: "really", c: "small"},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		ret := FooV(a)
		_ = ret
	}
}
