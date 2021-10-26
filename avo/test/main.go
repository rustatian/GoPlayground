package main

import (
	"github.com/rustatian/GoPlayground/avo/generated"
)

func main() {
	dig := &[4]uint32{1, 2, 3, 4}
	data := make([]byte, 100)
	generated.BlockScalar(dig, data)

	println(data)
}
