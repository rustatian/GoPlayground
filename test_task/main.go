package main

import (
	"github.com/ValeryPiashchynski/Sagas/zk"
)

func main() {
	addr := []string{"localhost:2181"}

	aa := zk.NewConfig()

	g, _ := zk.NewAion(addr, aa)
	stats := g.Close()

	print(stats.String())
}
