package main

import (
	_ "github.com/mmcloughlin/avo"
	_ "unsafe"
)

func sliceAdd(a []int, b int) []int {
	return append(a, b)
}
