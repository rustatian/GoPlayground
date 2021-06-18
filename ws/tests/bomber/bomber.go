package main

import (
	_ "embed"
	"flag"
	"net"
	"net/rpc"
	"time"

	"github.com/rustatian/GoPlayground/ws/tests/message"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
	"go.uber.org/ratelimit"
)

var addr = flag.String("addr", "localhost:6001", "amqp connection string")
var rate = flag.Uint64("r", 1, "filters rate per second")

func main() {
	flag.Parse()

	data := "iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii"

	m := makeMessage("join", "memory", []byte(data), "foo", "foo2")

	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		panic(err)
	}

	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	rl := ratelimit.New(int(*rate))
	_ = rl

	for j := 0; j < 1_000_000; j++ {
		for i := 0; i < 1_000_000; i++ {
			//_ = rl.Take()
			resp := &websocketsv1beta.Response{}
			err = client.Call("broadcast.Publish", m, resp)
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second * 5)
		}
	}
}
func makeMessage(command string, broker string, payload []byte, topics ...string) *websocketsv1beta.Request {
	m := &websocketsv1beta.Request{
		Messages: []*websocketsv1beta.Message{
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
