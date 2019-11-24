package main

import "testing"

//func BenchmarkPackWOptimization(b *testing.B) {
//	var s uint64 = 222
//	m := "some_long_stringggg"
//	b.ReportAllocs()
//	for n := 0; n < b.N; n++ {
//		data := PackWoOptimization(m, s)
//		_ = data
//	}
//}

func BenchmarkPackWoOptimization(b *testing.B) {
	var s uint64 = 222
	m := "some_long_stringggg"
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := PackWoOptimization(m, s)
		_ = data
	}
}