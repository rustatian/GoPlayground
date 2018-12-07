package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go foo(i)
		wg.Done()
	}
	wg.Wait()
}

func foo(i int) {
	defer fmt.Println(i)
}
