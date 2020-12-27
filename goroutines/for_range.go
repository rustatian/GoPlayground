package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{1, 2, 3}

	wg := &sync.WaitGroup{}
	wg.Add(3)
	for _, m := range a {
		aa := m
		go func() {
			fmt.Println(aa)
			wg.Done()
		}()
	}

	wg.Wait()
}
