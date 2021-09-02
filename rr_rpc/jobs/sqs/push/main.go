package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net"
	"net/rpc"

	"github.com/google/uuid"
	jobsv1beta "github.com/rustatian/GoPlayground/rr_rpc/jobs/v1beta"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
)

var addr = flag.String("addr", "localhost:6001", "")
var rate = flag.Uint64("r", 1, "filters rate per second")

func main() {
	flag.Parse()

	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func() {
			conn, err := net.Dial("tcp", *addr)
			if err != nil {
				fmt.Println(err)
				return
			}

			client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

			for j := 0; j < 100000; j++ {
				payloads := &jobsv1beta.PushRequest{
					Job: &jobsv1beta.Job{
						Job:     "fooooooo",
						Id:      uuid.NewString(),
						Payload: "test2",
						Headers: map[string]*jobsv1beta.HeaderValue{"test": {Value: []string{"hello"}}},
						Options: &jobsv1beta.Options{
							Priority: 1,
							Pipeline: "test-1",
						},
					},
				}

				resp := jobsv1beta.Empty{}
				err = client.Call("jobs.Push", payloads, &resp)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}

	<-ch
}
