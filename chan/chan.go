package main

import (
	"log"
	"sync"
)

var receiver chan int

func Setup() (receiver chan int) {
	receiver = make(chan int)
	return
}

//func Setup() (chan int) {
//  receiver = make(chan int)
//  return receiver
//}

func Launch(j int) {

	for i := 0; i < j; i++ {
		receiver <- i
	}

}

func main() {

	var wg sync.WaitGroup
	wg.Add(10)

	receiver = Setup()

	go func() {
		for r := range receiver {
			log.Println(r)
			wg.Done()
		}
	}()

	Launch(10)

	wg.Wait()

}
