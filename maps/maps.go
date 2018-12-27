package main

import "sync"

func main() {
	mmm := make(map[int]int, 10)
	wg := &sync.WaitGroup{}
	for i:=0; i < 10000; i ++ {
		wg.Add(1)
		go ch(mmm, wg)
	}
	wg.Wait()
}

func ch(m map[int]int, wg *sync.WaitGroup) {
	a := m[1]
	_ = a
	wg.Done()
}
