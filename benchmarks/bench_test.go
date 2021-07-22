package main

import (
	"crypto/rand"
	"testing"
)

//func Benchmark_One(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		ii := One()
//		_ = ii
//	}
//}
//
//func Benchmark_Two(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Two()
//	}
//}

func Benchmark_Ptr(b *testing.B) {
	b1 := make([]byte, 0, 1000)
	b2 := make([]byte, 0, 1000)

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

	for i := 0; i < b.N; i++ {
		l := fooPtr(&Payload{
			Context: b1,
			Body:    b2,
		})
		_ = l
	}
}

func Benchmark_Value(b *testing.B) {
	b1 := make([]byte, 0, 1000)
	b2 := make([]byte, 0, 1000)

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

	pld := Payload{
		Context: b1,
		Body:    b2,
	}

	for i := 0; i < b.N; i++ {
		l := fooVal(pld)
		_ = l
	}
}

func Benchmark_PtrVal(b *testing.B) {
	b1 := make([]byte, 0, 1000)
	b2 := make([]byte, 0, 1000)

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

	for i := 0; i < b.N; i++ {
		l := fooPtrVal(&Payload{
			Context: b1,
			Body:    b2,
		})
		_ = l
	}
}
// Benchmark_Pool_Echo_Batched-32    	  174765	      5723 ns/op	    7534 B/op	      20 allocs/op
/*
Benchmark_Pool_Echo_Batched-32    	  174646	      5826 ns/op	    7508 B/op	      20 allocs/op
Benchmark_Pool_Echo_Batched-32    	  175132	      5773 ns/op	    7534 B/op	      20 allocs/op

Benchmark_Pool_Echo_Batched-32    	  176272	      5755 ns/op	    7576 B/op	      21 allocs/op
Benchmark_Pool_Echo_Batched-32    	  215958	      5571 ns/op	    7426 B/op	      21 allocs/op


Benchmark_Pool_Echo-32    	   49076	     29926 ns/op	    8016 B/op	      20 allocs/op
Benchmark_Pool_Echo-32    	   47257	     30779 ns/op	    8047 B/op	      20 allocs/op
Benchmark_Pool_Echo-32    	   46737	     29440 ns/op	    8065 B/op	      20 allocs/op
Benchmark_Pool_Echo-32    	   51177	     29074 ns/op	    7981 B/op	      20 allocs/op
Benchmark_Pool_Echo-32    	   51764	     28319 ns/op	    8012 B/op	      20 allocs/op
Benchmark_Pool_Echo-32    	   54054	     30714 ns/op	    7987 B/op	      20 allocs/op

Benchmark_Pool_Echo-32    	   47936	     28679 ns/op	    7942 B/op	      19 allocs/op
Benchmark_Pool_Echo-32    	   49010	     29830 ns/op	    7970 B/op	      19 allocs/op
Benchmark_Pool_Echo-32    	   46771	     29031 ns/op	    8014 B/op	      19 allocs/op
Benchmark_Pool_Echo-32    	   47760	     30517 ns/op	    7955 B/op	      19 allocs/op
Benchmark_Pool_Echo-32    	   48148	     29816 ns/op	    7950 B/op	      19 allocs/op
Benchmark_Pool_Echo-32    	   52705	     29809 ns/op	    7979 B/op	      19 allocs/op
Benchmark_Pool_Echo-32    	   54374	     27776 ns/op	    7947 B/op	      19 allocs/op
 */
func Benchmark_ValuePtr(b *testing.B) {
	b1 := make([]byte, 0, 1000)
	b2 := make([]byte, 0, 1000)

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

	pld := Payload{
		Context: b1,
		Body:    b2,
	}

	for i := 0; i < b.N; i++ {
		l := fooValPtr(pld)
		_ = l
	}
}
