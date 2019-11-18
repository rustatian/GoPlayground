package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	err := errors.New("some error")

	e := fmt.Sprintf("%v", err)
	println(e)

	//rand.Seed(5)
	//wg := &sync.WaitGroup{}
	//
	//s := make([]int, 1e6)
	//for i := 0; i < len(s); i++ {
	//	s[i] = rand.Intn(1000)
	//}
	//cc(s, 0, wg)
}

func cc(data []int, numThreads int, wg *sync.WaitGroup) {
	if numThreads == 0 {
		numThreads = runtime.NumCPU()
	}

	//m := &sync.Mutex{}

	//result := make([]int, 0)

	begin := len(data) / numThreads

	// calculate remainder
	remainder := len(data) % numThreads
	println(remainder)

	for i := 0; i < numThreads; i++ {
		go func(index int) {
			dataCopy := make([]int, begin)
			if i == 0 {
				dataCopy = data[begin*i + 1 : begin*(i+1)]
			} else if remainder != 0 && i == numThreads {
				dataCopy = data[begin*i : begin*(i+1) + remainder]
			} else {
				dataCopy = data[begin*i : begin*(i+1)]
			}

			wg.Add(1)
			defer wg.Done()

			for j:=0; i < len(dataCopy); j++ {

			}

			//m.Lock()
			//result
			//m.Unlock()

		}(i)
	}

	wg.Wait()
}
