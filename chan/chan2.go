package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 5
	b := 6
	c := 7
	m := make(map[string]*int, 10)
	m["1"] = &a
	m["2"] = &b
	m["3"] = &c
	wg := &sync.WaitGroup{}

	for _, n := range m {
		nn := n
		wg.Add(1)
		go func() {
			fmt.Println(nn)
			wg.Done()
		}()
	}

	wg.Wait()
}


