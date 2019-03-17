package util_test

import (
	"github.com/ValeryPiashchynski/GoPlayground/time"
	"testing"
)

func TestFTime_UnixNano(t *testing.T) {
	// ft := NewFastTime(time.Millisecond)
	// defer ft.Stop()

	// log.Println(ft.UnixNano())
	// log.Println(time.Now().UnixNano())
}

func Benchmark_FastTime(b *testing.B) {
	b.ReportAllocs()
	ft := util.Now()

	var v int64
	for n := 0; n < b.N; n++ {
		v = ft.UnixNano()
		_ = v
	}

	if v == 0 {
		panic("invalid")
	}
}

//
//func Benchmark_Time(b *testing.B) {
//	b.ReportAllocs()
//	v := time.Now()
//	var val int64 = 4
//
//	tmp := int64(4)
//	valAddr := &tmp
//
//	atomic.StoreInt64(valAddr, val)
//
//	for n := 0; n < b.N; n++ {
//		va := atomic.LoadInt64(valAddr)
//		_ = va
//		v.UnixNano()
//	}
//
//}

//Benchmark_FastTime-8   	2000000000	         1.48 ns/op	       0 B/op	       0 allocs/op
//Benchmark_Time-8       	20000000	        87.5 ns/op
