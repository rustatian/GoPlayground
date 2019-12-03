package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt64 == (int64(^uint64(0) >> 1)))
}


//func returnMaxIntOptimized() int64 {
//	return int64(^uint64(0) >> 1)
//}
//
//func returnMaxIntNotOptimized() int64 {
//	return math.MaxInt64
//}
