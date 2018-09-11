package funtest

import (
	"github.com/ValeryPiashchynski/GoPlayground/lowlevelfun/funStrings"
	"testing"
)

func Benchmark_Fun(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		funStrings.FromInterfaceToBytes([]byte{1, 2, 3})
	}
}

func Benchmark_NonFun(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		funStrings.NonFunConvert("some text for parse")
	}
}
