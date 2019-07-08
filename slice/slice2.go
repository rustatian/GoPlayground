package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make([]byte, 10000)
	wg := &sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i <= 50; i++ {
		go writeData(a, wg)
		go writeData2(a, wg)
	}

	wg.Wait()

	fmt.Print(a)
}

func writeData(data []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	data[1] = byte('1')
}

func writeData2(data []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	data[2] = byte('2')
}

func DoTheWork(data []interface{}, numGoroutines int) {

}