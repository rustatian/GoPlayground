package main

import "cmp"

func main() {
	foo := 0
	print(cmp.Or(foo, -1))
}
