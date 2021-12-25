package main

import (
	"strings"
)

func main() {
	a := "foo\n\r"
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, "\r", "")
	print(a)
}
