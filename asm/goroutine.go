package main

import "fmt"

func main() {
	var a, b, c int = 1, 2, 3
	go someFunc(a, b, c)
}

func someFunc(a, b, c int) {
	fmt.Print(a + b + c)
}
