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

	//payload := &jobsv1beta.Request{
	//	Job: &jobsv1beta.Job{
	//		Job:     "test",
	//		Payload: "helllllllllllllllllllloooooooooooooooooooo",
	//		Options: &jobsv1beta.Options{
	//			Priority:   100,
	//			Pipeline:   "test-local",
	//			Delay:      0,
	//			Attempts:   0,
	//			RetryDelay: 0,
	//			Timeout:    0,
	//		},
	//	},
	//}

	payloads := &jobsv1beta.PushBatchRequest{
		Jobs: []*jobsv1beta.Job{
			{
				Job:     "test1",
				Id:      "1",
				Payload: "test1-fffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 11,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test2",
				Id:      "1",
				Payload: "test2-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 10,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test3",
				Id:      "1",
				Payload: "test3-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 12,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test1",
				Id:      "1",
				Payload: "test1-fffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 11,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test2",
				Id:      "1",
				Payload: "test2-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 10,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test3",
				Id:      "1",
				Payload: "test3-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 12,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test1",
				Id:      "1",
				Payload: "test1-fffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 11,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test2",
				Id:      "1",
				Payload: "test2-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 10,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test3",
				Id:      "1",
				Payload: "test3-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 12,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test1",
				Id:      "1",
				Payload: "test1-fffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 11,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test2",
				Id:      "1",
				Payload: "test2-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 10,
					Pipeline: "test-local",
				},
			},
			{
				Job:     "test3",
				Id:      "1",
				Payload: "test3-ffffffffffffffff",
				Options: &jobsv1beta.Options{
					Priority: 12,
					Pipeline: "test-local",
				},
			},
		},
	}

	ch := make(chan struct{})
	for i := 0; i < 20; i++ {
		go func() {
			conn, err := net.Dial("unix", *addr)
			if err != nil {
				panic(err)
			}

			client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

			for j := 0; j < 10000; j++ {
				resp := jobsv1beta.EmptyResponse{}
				err = client.Call("jobs.PushBatch", payloads, &resp)
				if err != nil {
					panic(err)
				}
			}
		}()
	}

	<-ch
}
