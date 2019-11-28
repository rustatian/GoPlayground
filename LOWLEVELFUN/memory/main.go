package main

import "runtime"

func main() {
	//f, err := elf.Open("/Users/0xdev/Projects/repo/GoPlayground/LOWLEVELFUN/memory/main")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, section := range f.Sections {
	//	log.Println(section)
	//}
	a := make([]int, 100)
	runtime.GC()
	runtime.KeepAlive(a)
}
