package main

import (
	"sync"
)

func main() {
	c := make(chan int, 100)
	wg := &sync.WaitGroup{}

	foo(c, wg)



}

func foo(c chan int, wg *sync.WaitGroup) {

}

