package main

import "fmt"

func main() {
	a := make(map[int]int, 6)
	a[1] = 2
	changeMap(a)
	fmt.Print(a[1])
}

func changeMap(m map[int]int) {
	m[1] = 23
}
