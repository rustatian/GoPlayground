package main

import (
	"debug/elf"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func main() {
	//00000000004aa450
	testname := "bitbucket.org/inturnco/productservice/search.service.FetchProductIDList"
	println(testname)
	path := "/home/valery/Projects/repo/GoPlayground/LOWLEVELFUN/elf/main"
	elffile, err := readELFfile(path)
	if err != nil {
		panic(err)
	}

	findInELF(elffile)

}


func readELFfile(path string) (*elf.File, error) {
	f, err := os.Open(path)
	// defer f.Close()
	if err != nil {
		return nil, err
	}

	return elf.NewFile(f)
}

// from https://github.com/golang/go/blob/dcd3b2c173b77d93be1c391e3b5f932e0779fb1f/src/reflect/makefunc.go#L56-L60
func getFunctionAddress(func_name interface{}) uintptr {
	type dummy struct {
		typ uintptr
		value *uintptr
	}
	e := (*dummy)(unsafe.Pointer(&func_name))
	return *e.value
}

func findInELF(elffile *elf.File) {
	s, err := elffile.Symbols()
	if err != nil {
		panic(err)
	}
	//println(elffile.Data.String())

	sectionnn := elffile.Section(".text")
	println(sectionnn.SectionHeader.Info)

	//for _, v := range elffile.Sections {
	//	println(v.Name)
	//}
	//
	for _, symbol := range s {
		//00000000004aa450
		if symbol.Value == uint64(0x4f1d30) {
			println("FOUND")

			code := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
				Data: uintptr(0x4aa450),
				Len:  int(symbol.Size),
				Cap:  int(symbol.Size),
			}))

			fmt.Printf("% x\n", code)
		}
	}
}
