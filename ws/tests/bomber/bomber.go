package main

import (
	_ "embed"
	"flag"
	"net"
	"net/rpc"

	"github.com/48d90782/GoPlayground/ws/tests/message"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
	"go.uber.org/ratelimit"
)

var addr = flag.String("addr", "localhost:6001", "amqp connection string")
var rate = flag.Uint64("r", 1000000, "filters rate per second")

func main() {
	flag.Parse()

	data := "hello bbbbbbbbbbbbbbbeeeeeeeeeee"

	m := makeMessage("join", "memory", []byte(data), "foo", "foo2")

	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	rl := ratelimit.New(int(*rate))

	for j := 0; j < 1_000_000; j++ {
		for i := 0; i < 1_000_000; i++ {
			_ = rl.Take()
			var ret bool
			err = client.Call("websockets.Publish", m, &ret)
			if err != nil {
				panic(err)
			}
		}
	}
}
func makeMessage(command string, broker string, payload []byte, topics ...string) *message.Messages {
	m := &message.Messages{
		Messages: []*message.Message{
			{
				Topics:  topics,
				Command: command,
				Broker:  broker,
				Payload: payload,
			},
		},
	}

	return m
}
