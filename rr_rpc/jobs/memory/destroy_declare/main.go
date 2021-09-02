package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/google/uuid"
	jobsv1beta "github.com/rustatian/GoPlayground/rr_rpc/jobs/v1beta"
	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
)

const (
	push    string = "jobs.Push"
	pause   string = "jobs.Pause"
	destroy string = "jobs.Destroy"
	declare string = "jobs.Declare"
	resume  string = "jobs.Resume"
	stat    string = "jobs.Stat"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))
	for kk := 0; kk < 10; kk++ {
		for i := 0; i < 5; i++ {
			go func() {
				for j := 0; j < 500; j++ {
					n := uuid.NewString()
					declareMemoryPipe(client, n)
					startPipelines(client, n)
					push100(client, n)
					pausePipelines(client, n)
					destroyPipelines(client, n)
				}
				wg.Done()
			}()
		}
		time.Sleep(time.Second * 5)
	}
	wg.Wait()
}

func startPipelines(client *rpc.Client, pipes ...string) {
	pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

	for i := 0; i < len(pipes); i++ {
		pipe.GetPipelines()[i] = pipes[i]
	}

	er := &jobsv1beta.Empty{}
	err := client.Call(resume, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}

func pausePipelines(client *rpc.Client, pipes ...string) {
	pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

	for i := 0; i < len(pipes); i++ {
		pipe.GetPipelines()[i] = pipes[i]
	}

	er := &jobsv1beta.Empty{}
	err := client.Call(pause, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}

func push100(client *rpc.Client, pipe string) {
	for j := 0; j < 1000; j++ {
		payloads := &jobsv1beta.PushRequest{
			Job: &jobsv1beta.Job{
				Job:     "fooooooo",
				Id:      uuid.NewString(),
				Payload: "test2",
				Headers: map[string]*jobsv1beta.HeaderValue{"test": {Value: []string{"hello"}}},
				Options: &jobsv1beta.Options{
					Priority: 1,
					Pipeline: pipe,
				},
			},
		}

		resp := jobsv1beta.Empty{}
		err := client.Call(push, payloads, &resp)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func destroyPipelines(client *rpc.Client, pipes ...string) {
	pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

	for i := 0; i < len(pipes); i++ {
		pipe.GetPipelines()[i] = pipes[i]
	}

	er := &jobsv1beta.Empty{}
	err := client.Call(destroy, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}

func declareMemoryPipe(client *rpc.Client, p string) {
	pipe := &jobsv1beta.DeclareRequest{Pipeline: map[string]string{
		"driver":   "memory",
		"name":     p,
		"prefetch": "10000",
	}}

	er := &jobsv1beta.Empty{}
	err := client.Call(declare, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}
