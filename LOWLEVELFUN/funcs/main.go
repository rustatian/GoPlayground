package main

import (
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(ii int) {
			print(ii)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
