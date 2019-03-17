package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	FindFuncWithName("slice.xy")

}

//go:linkname Firstmoduledata runtime.firstmoduledata
var Firstmoduledata Moduledata

func FindFuncWithName(name string) (uintptr, error) {
	for moduleData := &Firstmoduledata; moduleData != nil; moduleData = moduleData.next {
		for _, ftab := range moduleData.ftab {
			f := (*runtime.Func)(unsafe.Pointer(&moduleData.pclntable[ftab.funcoff]))
			if f.Name() == name {
				return f.Entry(), nil
			}
		}
	}
	return 0, fmt.Errorf("Invalid function name: %s", name)
}

type Moduledata struct {
	pclntable    []byte
	ftab         []Functab
	filetab      []uint32
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	end, gcdata, gcbss    uintptr

	// Original type was []*_type
	typelinks []interface{}

	modulename string
	// Original type was []modulehash
	modulehashes []interface{}

	gcdatamask, gcbssmask Bitvector

	next *Moduledata
}

type Functab struct {
	entry   uintptr
	funcoff uintptr
}

type Bitvector struct {
	n        int32 // # of bits
	bytedata *uint8
}
