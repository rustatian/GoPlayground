package main

import (
	_ "embed"
	"fmt"
	"net"
	"net/rpc"

	goridgeRpc "github.com/roadrunner-server/goridge/v3/pkg/rpc"
)

func main() {
	conn, err := net.Dial("tcp", "100.100.109.23:6001")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	data := make([]string, 0, 100)
	err = client.Call("temporal.GetActivityNames", false, &data)
	fmt.Println(data)
}
