package main

import "fmt"

var str1 = "123"
var str2 = "123"

func main() {
	if &str1 != &str2 {
		fmt.Print("Whooohoo!")
	}
}
