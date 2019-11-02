package inter

import (
	"testing"
)

func Benchmark_Iface_Val(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ii := InterfaceFuncVal(someVarVal)
		_ = ii
	}
}

func Benchmark_Iface_Ptr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ii := InterfaceFuncPtr(someVarPtr)
		_ = ii
	}
}

//go:noinline
func InterfaceFuncPtr(i interface{}) interface{} {
	return i
}

//go:noinline
func InterfaceFuncVal(i interface{}) interface{} {
	return i
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

var someVarPtr = &foo{
	name:   "fasdfasdf",
	name2:  "adfasdf",
	name3:  "asdfawer",
	name4:  "adsfasdfas",
	name5:  "asdfewrqae",
	name6:  "faerwefas",
	name7:  "faserwfwaee",
	name8:  "fdsageawegawwg",
	name9:  "fweff34fsdfas",
	name10: "fasgharhdshdhewr344",
	name11: "asdf432fdsagas",
	name12: "f3w4f3sdafaw",
	name13: "fawf3443fsafgagha",
	name14: "gargrwgehsdgnhfdgbda",
	name15: "fsadgasgneworghaoisgnoisdufhbphgviuashdkjasdlkvbnasdkjvnb",
}
var someVarVal = foo{
	name:   "fasdfasdf",
	name2:  "adfasdf",
	name3:  "asdfawer",
	name4:  "adsfasdfas",
	name5:  "asdfewrqae",
	name6:  "faerwefas",
	name7:  "faserwfwaee",
	name8:  "fdsageawegawwg",
	name9:  "fweff34fsdfas",
	name10: "fasgharhdshdhewr344",
	name11: "asdf432fdsagas",
	name12: "f3w4f3sdafaw",
	name13: "fawf3443fsafgagha",
	name14: "gargrwgehsdgnhfdgbda",
	name15: "fsadgasgneworghaoisgnoisdufhbphgviuashdkjasdlkvbnasdkjvnb",
}
