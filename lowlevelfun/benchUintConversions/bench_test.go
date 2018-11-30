package main

import (
	"encoding/binary"
	"testing"
)

//go:noinline
func convertFun(id [8]byte) uint64 {
	var u64 uint64
	u64 |= uint64(id[7]&0xFF) << 0
	u64 |= uint64(id[6]&0xFF) << 8
	u64 |= uint64(id[5]&0xFF) << 16
	u64 |= uint64(id[4]&0xFF) << 24
	u64 |= uint64(id[3]&0xFF) << 32
	u64 |= uint64(id[2]&0xFF) << 40
	u64 |= uint64(id[1]&0xFF) << 48
	u64 |= uint64(id[0]&0xFF) << 56
	return u64
}

//go:noinline
func convertNotFun(id []byte) uint64 {
	return binary.BigEndian.Uint64(id)
}

func Benchmark_FunConvert(b *testing.B) {
	var byteSlice = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		convertFun(byteSlice)
	}
}

func Benchmark_NotFunConvert(b *testing.B) {
	var byteSlice = []byte{1, 2, 3, 4, 5, 6, 7, 8}
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		convertNotFun(byteSlice)
	}
}
