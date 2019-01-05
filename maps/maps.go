package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i:=0; i < 10; i ++ {
		wg.Add(2)
		go ch("ch", wg)
		go ch("ch2", wg)
	}
	wg.Wait()
}

func ch(a string, wg *sync.WaitGroup) {
	fmt.Println(a)
	wg.Done()
}
