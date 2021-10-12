package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	// Create JetStream Context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		panic(err)
	}

	// Simple Stream Publisher
	_, err = js.Publish("ORDERS.scratch", []byte("hello"))
	if err != nil {
		panic(err)
	}

	// Simple Async Stream Publisher
	for i := 0; i < 500; i++ {
		_, err = js.PublishAsync("ORDERS.scratch", []byte("hello"))
		if err != nil {
			panic(err)
		}
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// Simple Async Ephemeral Consumer
	_, err = js.Subscribe("ORDERS.*", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})
	if err != nil {
		panic(err)
	}

	// Simple Sync Durable Consumer (optional SubOpts at the end)
	sub, err := js.SubscribeSync("ORDERS.*", nats.Durable("MONITOR"), nats.MaxDeliver(3))
	m, err := sub.NextMsg(time.Second * 2)
	if err != nil {
		panic(err)
	}

	println(m.Data)

	// Simple Pull Consumer
	sub, err = js.PullSubscribe("ORDERS.*", "MONITOR")
	if err != nil {
		panic(err)
	}
	msgs, err := sub.Fetch(10)
	if err != nil {
		panic(err)
	}

	println(msgs)

	// Unsubscribe
	err = sub.Unsubscribe()
	if err != nil {
		panic(err)
	}

	// Drain
	err = sub.Drain()
	if err != nil {
		panic(err)
	}
}
