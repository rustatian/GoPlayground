package main

import (
	"os"
	"strconv"
	"time"
)

var max int

func main() {
	max, _ = strconv.Atoi(os.Args[1])
	n := 0
	for {
		time.Sleep(time.Second / 100)
		n += 1
		if n >= max {
			return
		}
	}
}