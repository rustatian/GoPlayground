package main

import (
	endure "github.com/roadrunner-server/endure/pkg/container"
	"github.com/rustatian/GoPlayground/endure/p1"
	"github.com/rustatian/GoPlayground/endure/p2"
)

func main() {
	c, err := endure.NewContainer(nil)
	if err != nil {
		panic(err)
	}

	err = c.RegisterAll(
		&p1.Plugin{},
		&p2.Plugin{},
	)

	if err != nil {
		panic(err)
	}

	err = c.Init()
	if err != nil {
		panic(err)
	}

	ch, err := c.Serve()
	if err != nil {
		panic(err)
	}

	for c := range ch {
		println(c.Error)
	}
}
