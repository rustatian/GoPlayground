package main

import (
	//"fmt"
	"runtime/debug"
)

func main() {
	//x := "fdfsdf"

	panic("dfasdfsdf")

	defer func() {
		if p := recover(); p != nil {
			debug.PrintStack()
		}
	}()
}
