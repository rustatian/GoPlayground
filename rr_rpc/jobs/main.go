package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net"
	"net/rpc"

	jobsv1beta "github.com/rustatian/GoPlayground/rr_rpc/jobs/v1beta"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
)

var addr = flag.String("addr", "/tmp/rr.sock", "")
var rate = flag.Uint64("r", 1, "filters rate per second")

func main() {
	flag.Parse()

	payloads := &jobsv1beta.PushBatchRequest{
		Jobs: []*jobsv1beta.Job{
			{
				Job:     "test1",
				Id:      "1",
				Payload: "priority1",
				Options: &jobsv1beta.Options{
					Priority: 1,
					Pipeline: "test-1",
					Delay:    5,
				},
			},
			{
				Job:     "test2-",
				Id:      "2",
				Payload: "amqp-2",
				Headers: map[string]*jobsv1beta.HeaderValue{"test": {Value: []string{"hello"}}},
				Options: &jobsv1beta.Options{
					Priority: 1,
					Pipeline: "test-2-amqp",
				},
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

			for j := 0; j < 9999999; j++ {
				resp := jobsv1beta.EmptyResponse{}
				err = client.Call("jobs.PushBatch", payloads, &resp)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	<-ch
}
