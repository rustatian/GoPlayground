package main

import (
	"fmt"
	"time"
)

func main() {
	foo()

	time.Sleep(time.Second * 5)
}
func foo() int {
	donec := make(chan struct{})
	defer close(donec)
	go func() {
		ticker := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-ticker.C:
				fmt.Println("RUNNING C")
			case <-donec:
				fmt.Println("DONE")
				ticker.Stop()
				return
			}
		}
	}()
	return fooo()
}

func fooo() int {
	time.Sleep(time.Second * 4)
	return 5
}
