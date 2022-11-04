package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("/tmp/foo.err", os.O_RDWR, 0)
	if err != nil {
		panic(err)
	}

	off := int64(0)

	for {
		b := make([]byte, 10)
		n, err := f.ReadAt(b, off)
		if err != nil {
			panic(err)
		}
		off += int64(n)
		fmt.Println(string(b))
		time.Sleep(time.Second)
	}
}
