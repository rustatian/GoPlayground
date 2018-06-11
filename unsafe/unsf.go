package main

import "unsafe"

func main() {
	const aa = 5
	ab := aa
	ab = 6
	a := unsafe.Pointer(&ab)

	println(aa)
	println(ab)

}
