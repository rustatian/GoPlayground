package main

import "fmt"
import "runtime"

func main() {
	c := make(chan int)
	go func() {
		defer func() {c <- 1}()
		defer fmt.Println("Go")
		func() {
			defer fmt.Println("C")
			runtime.Goexit()
		}()
		fmt.Println("Java")
	}()
	<-c
}