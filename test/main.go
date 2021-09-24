package main

import (
	"time"
)

func main() {

}

//go:noinline
func sleep() {
	time.Sleep(time.Millisecond * 10)
}

//go:noinline
func block() {
	tt := time.After(time.Millisecond * 10)
	select {
	case <-tt:

	}
}
