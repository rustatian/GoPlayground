package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5}
	fmt.Println(&a[2])
	fmt.Println(&a[3])
}
