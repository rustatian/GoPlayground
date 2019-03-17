package main

import (
	tm "github.com/buger/goterm"
	"runtime"
)

func main() {
	var mem runtime.MemStats

	go func() {
		for {
			runtime.ReadMemStats(&mem)
		}
	}()

	a := make(map[int]string)
	for i := 0; i < 100000; i++ {
		tm.MoveCursor(1, 1)
		a[i] = "Lorem input"
		tm.Printf("Memory alloc: %d | Total alloc: %d | Heap alloc: %d | HeapSys: %d, Number of GCs: %d", mem.HeapInuse, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys, mem.NumGC)
		tm.Flush()
	}
}
