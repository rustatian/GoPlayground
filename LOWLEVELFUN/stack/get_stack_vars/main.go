package main

import (
	"testing"
)

var someVariable = 1000
var aaaa = &someVariable

func some_test(t *testing.T) {

}

func main() {

	println(aaaa)
	println(someVariable)

}
