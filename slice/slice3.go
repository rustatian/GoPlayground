package main

import "fmt"

func main() {
	ints := []int{1,2,3,4,5,6}
	var out []*int

	for i, _ := range ints {
		fmt.Println(&ints[i])
		out = append(out, &ints[i])
		fmt.Println(&out[i])
		fmt.Println("8=======================0")
	}
}
