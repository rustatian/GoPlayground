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
	for i := 0; i < 1; i++ {
		go func() {
			for j := 0; j < 1; j++ {
				n := uuid.NewString()
				declareAMQPPipe(n)
				startPipelines(n)
				push100(n)
				pausePipelines(n)
				time.Sleep(time.Second)
				destroyPipelines(n)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func push100(pipe string) {
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	for j := 0; j < 100; j++ {
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

func startPipelines(pipes ...string) {
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

	for i := 0; i < len(pipes); i++ {
		pipe.GetPipelines()[i] = pipes[i]
	}

	er := &jobsv1beta.Empty{}
	err = client.Call(resume, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}

func pausePipelines(pipes ...string) {
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

	for i := 0; i < len(pipes); i++ {
		pipe.GetPipelines()[i] = pipes[i]
	}

	er := &jobsv1beta.Empty{}
	err = client.Call(pause, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}

func destroyPipelines(pipes ...string) {
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	pipe := &jobsv1beta.Pipelines{Pipelines: make([]string, len(pipes))}

	for i := 0; i < len(pipes); i++ {
		pipe.GetPipelines()[i] = pipes[i]
	}

	er := &jobsv1beta.Empty{}
	err = client.Call(destroy, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}

func declareAMQPPipe(p string) {
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		log.Fatal(err)
	}
	client := rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	pipe := &jobsv1beta.DeclareRequest{Pipeline: map[string]string{
		"driver":          "amqp",
		"name":            p,
		"routing_key":     "test-3",
		"queue":           "default",
		"exchange_type":   "direct",
		"exchange":        "amqp.default",
		"prefetch":        "100",
		"priority":        "3",
		"exclusive":       "false",
		"multiple_ask":    "false",
		"requeue_on_fail": "false",
	}}

	er := &jobsv1beta.Empty{}
	err = client.Call(declare, pipe, er)
	if err != nil {
		log.Fatal(err)
	}
}
