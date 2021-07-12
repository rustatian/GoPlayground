package main

import "sync/atomic"

func main() {
	a := 10
	val := atomic.Value{}

	val.Store(a)

	b := val.Load().(int)
	println(b)

	c := val.Load().(int)
	println(c)
}
