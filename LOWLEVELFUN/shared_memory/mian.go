package main

import (
	"fmt"

	"github.com/roadrunner-server/goridge/v3/pkg/shared_memory/posix"
)

func main() {
	seg, err := posix.AttachToShmSegment(6815800, 1024, 0666)
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1024)
	err = seg.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
