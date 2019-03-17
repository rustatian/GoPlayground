package main

import (
	"fmt"
	"sync"
)

func main() {
	tokens := make(chan struct{}, 10)
	wg := sync.WaitGroup{}
	lines := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := range lines {
		i := i
		wg.Add(1)
		tokens <- struct{}{}

		go func() {
			lines[i] = replaceLink(lines[i])
			<-tokens
			wg.Done()
		}()
	}

	wg.Wait()

	for _, line := range lines {
		fmt.Println(line)
	}
}
func replaceLink(i int) int {
	return 2
}
