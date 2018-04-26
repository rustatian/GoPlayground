package main

import (
	"github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		println(err)
	}

	// Simple Publisher
	nc.Publish("hash", []byte(`{"password":"20574178"}`))

	nc.Close()
}
