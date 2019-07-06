package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string, 10)
	if v := m["1"]; v != "" {
		fmt.Println("DATA")
		fmt.Println(v)
	}

}

func ch(a string, wg *sync.WaitGroup) {

}
