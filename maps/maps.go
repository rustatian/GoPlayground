package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[*int]string)
	m[nil] = "0"
	fmt.Println(m[nil])
	sort.Sort()
}

func changeMap(m map[int]int) {
	m[1] = 23
}
