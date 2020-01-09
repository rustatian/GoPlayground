package main

import "fmt"

func SuperPuperTestFuction(name string) {
	fmt.Print(name)
}

func main() {
	c := 3
	d := 5
	a:= make([]int, 0, c-d)
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	print(a)
}

