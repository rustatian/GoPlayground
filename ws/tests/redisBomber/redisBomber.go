package main

import (
	_ "embed"
	"flag"
	"net"
	"net/rpc"

	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
	"go.uber.org/ratelimit"
)

var addr = flag.String("addr", "amqp://appuser:eLVWJfhHjmMXwiLoMNVPY7cKVWO@127.0.0.1:5672/", "amqp connection string")
var rate = flag.Uint64("r", 1, "filters rate per second")

type Msg struct {
	// Topic message been pushed into.
	Topics_ []string `json:"topic"`

	// Command (join, leave, headers)
	Command_ string `json:"command"`

	// Broker (redis, memory)
	Broker_ string `json:"broker"`

	// Payload to be broadcasted
	Payload_ []byte `json:"payload"`
}

func main() {
	flag.Parse()

	data := "hello bbbbbbbbbbbbbbbeeeeeeeeeee"
	m := []*Msg{
		{
			Topics_:  []string{"foo", "foo2"},
			Command_: "join",
			Broker_:  "memory",
			Payload_: []byte(data),
		},
	}

	//d, err := json.Marshal(m)
	//if err != nil {
	//	panic(err)
	//}

	conn, err := net.Dial("tcp", "127.0.0.1:6001")
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
