package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	aaa := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(aaa); i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()

			fmt.Println(ii)
		}(i) //change(aaa[i], &wg)
	}
	wg.Wait()
}

func change(sl int, group *sync.WaitGroup) {
	fmt.Println(sl)
	group.Done()
}
