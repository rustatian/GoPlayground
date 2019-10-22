package main

import (
	"testing"
	"time"
)

func TestFTime_UnixNano(t *testing.T) {
	// ft := NewFastTime(time.Millisecond)
	// defer ft.Stop()

	// log.Println(ft.UnixNano())
	// log.Println(time.Now().UnixNano())
}

func Benchmark_FastTime(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Now().UnixNano()
	}
}

//
func Benchmark_Time(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		time.Now().UnixNano()
	}

}

//Benchmark_FastTime-8   	2000000000	         1.48 ns/op	       0 B/op	       0 allocs/op
//Benchmark_Time-8       	20000000	        87.5 ns/op
