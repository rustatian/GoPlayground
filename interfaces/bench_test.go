package inter

import (
	"testing"
)

func Benchmark_Iface_Val(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		InterfaceFuncVal(someVarVal)
	}
}

func Benchmark_Iface_Ptr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		InterfaceFuncPtr(someVarPtr)
	}
}

//go:noinline
func InterfaceFuncPtr(i interface{}) {
	i = nil
}

//go:noinline
func InterfaceFuncVal(i interface{}) {
	i = nil
}

type foo struct {
	name   string
	name2  string
	name3  string
	name4  string
	name5  string
	name6  string
	name7  string
	name8  string
	name9  string
	name10 string
	name11 string
	name12 string
	name13 string
	name14 string
	name15 string
}

var someVarPtr = &foo{}
var someVarVal = foo{}
