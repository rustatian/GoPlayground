package main

import "fmt"

func main() {
	cc := testtt()

	for v := range cc {
		fmt.Print(v)
	}

	fmt.Println("END")

}

func testtt() chan int {
	c := make(chan int, )
	go func() {
		for i := 0; i < 50; i ++ {
			c <- 1
		}
	}()
	return c
}
