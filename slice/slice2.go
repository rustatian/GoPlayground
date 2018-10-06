package main

import "fmt"

func main() {
	a := make([]byte, 0)

	for i := 0; i < 10; i++ {
		a = append(a, byte(i))
	}

	fmt.Print(a)
}
