package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)
		_, _ = fmt.Fprintf(os.Stderr, "Hello %d", i)
	}
}
