package funtest

import (
	"testing"
)

func Benchmark_Fun(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		FromInterfaceTo("some text for parse")
	}
}

func Benchmark_NonFun(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		funStrings.NonFunConvert("some text for parse")
	}
}
