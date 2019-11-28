package main

import (
	"fmt"
	_ "unsafe"
)

//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64)

func main() {
	a, _, _ := time_now()
	fmt.Print(a)
}
