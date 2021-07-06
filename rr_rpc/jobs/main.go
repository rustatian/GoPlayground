package main

import (
	_ "embed"
	"flag"
	"net"
	"net/rpc"

	jobsv1beta "github.com/rustatian/GoPlayground/rr_rpc/jobs/v1beta"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
)

var addr = flag.String("addr", "/home/valery/Downloads/rr.sock", "")
var rate = flag.Uint64("r", 1, "filters rate per second")

func main() {
	flag.Parse()

	payload := &jobsv1beta.Request{
		Job: &jobsv1beta.Job{
			Job:     "test",
			Payload: "helllllllllllllllllllloooooooooooooooooooo",
			Options: &jobsv1beta.Options{
				Priority:   100,
				Pipeline:   "test-local",
				Delay:      0,
				Attempts:   0,
				RetryDelay: 0,
				Timeout:    0,
			},
		},
	}

	ch := make(chan struct{})
	for i := 0; i < 1000; i++ {
		go func() {
			conn, err := net.Dial("unix", *addr)
			if err != nil {
				panic(err)
			}

			client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

			for j := 0; j < 20000000; j++ {
				resp := jobsv1beta.Response{}
				err = client.Call("jobs.Push", payload, &resp)
				if err != nil {
					panic(err)
				}
			}
		}()
	}

	<-ch
}
