package main

import (
	"fmt"
	"os"
)

func main() {
	for _, b := range os.Environ() {
		fmt.Println(b)
	}
}


