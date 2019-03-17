package main

import "unsafe"

func main() {
	const aa = 5
	ab := aa
	ab = 6
	_ = unsafe.Pointer(&ab)

	println(aa)
	println(ab)

}
