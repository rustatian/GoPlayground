package main

import (
	"fmt"
)

//go:noinline
func Square(a, b string) string {
	const foo = "some const string %s %s"

	return fmt.Sprintf(foo, a, b)
}

//go:noinline
func Square2(a, b string) string {
	const foo = "som3e const string %s %s"

	return fmt.Sprintf(foo, a, b)
}

func main() {
	_ = Square("3", "4")
	_ = Square2("1", "2")
}
