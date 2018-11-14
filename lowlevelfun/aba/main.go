package main

import "fmt"

func main() {
	a := 1
	b := 2
	a, b, a = b, a, b
	fmt.Print(a, b)
}
