package main

import (
	"github.com/roadrunner-server/roadrunner/v2/lib"
)

func main() {
	rr, err := lib.NewRR("/home/valery/projects/opensource/github/spiral/roadrunner/.rr.yaml", nil, nil)
	if err != nil {
		panic(err)
	}

	_ = rr
	println("foo")
}
