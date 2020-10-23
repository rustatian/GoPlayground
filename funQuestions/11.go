package main

import "fmt"

var numRun = 0

func init() {
	main()
	numRun +=1
}

func main() {
	if numRun < 1 {
		fmt.Println("fuck you YAUHENI")
		return
	}
	fmt.Println("fuck you VALERY")
}
