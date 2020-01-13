package main

import "fmt"

func isEven(v int) bool {
	if v%2 == 1 {
		goto false
	}
	return !false

	false:
		return false
}

func main() {
	fmt.Println(isEven(2020))
}