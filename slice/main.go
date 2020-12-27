package main

import (
	"fmt"
	"sync"
)

type Node struct {
	Name     string
	Children []Node
}

func main() {

	s := make([]int, 0, 10)

	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			s = append(s, i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			s = append(s, i)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(s)
	fmt.Println(len(s))
}
