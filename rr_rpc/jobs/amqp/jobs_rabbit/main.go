package main

import (
	"flag"
	"log"

	"github.com/google/uuid"
	json "github.com/json-iterator/go"
	"github.com/streadway/amqp"
)

var addr = flag.String("addr", "amqp://guest:guest@localhost:5672/", "amqp connection string")

func main() {
	flag.Parse()

	conn, err := amqp.Dial(*addr)
	if err != nil {
		log.Fatal(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan struct{})

	for j := 0; j < 100; j++ {
		go func() {
			for i := 0; i < 1_000_000; i++ {
				msg, err := pack("212", &Job{
					Job:     "super-job",
					Ident:   uuid.NewString(),
					Payload: "asdfkasjdf;laskjdflkajsdfl;kjas;ldkfjlksjdlfkasjdf;lkajs;dkjfkl;asjdfl",
					Headers: nil,
					Options: &Options{
						Priority: 1,
						Pipeline: "test-1",
					},
				})

				if err != nil {
					log.Fatal(err)
				}
				err = channel.Publish(
					"default",
					"test",
					false,
					false,
					amqp.Publishing{
						Headers:     msg,
						ContentType: "application/octet-stream",
						Body:        []byte("asdfal;sjf;lasjdf;lkjaal;sdfj;lasdfasdfjlasjdf;ljasd;fljasl;djf;lsajdfkajsdl;fk"),
					},
				)
				if err != nil {
					log.Fatal(err)
				}
			}
		}()
	}

	<-ch
}

// constant keys to pack/unpack messages from different drivers
const (
	RRID       string = "rr_id"
	RRJob      string = "rr_job"
	RRHeaders  string = "rr_headers"
	RRPipeline string = "rr_pipeline"
	RRDelay    string = "rr_delay"
	RRPriority string = "rr_priority"
)

// pack job metadata into headers
func pack(id string, j *Job) (amqp.Table, error) {
	headers, err := json.Marshal(j.Headers)
	if err != nil {
		return nil, err
	}
	return amqp.Table{
		RRID:       id,
		RRJob:      j.Job,
		RRPipeline: j.Options.Pipeline,
		RRHeaders:  headers,
		RRDelay:    j.Options.Delay,
		RRPriority: j.Options.Priority,
	}, nil
}

// Job carries information about single job.
type Job struct {
	// Job contains name of job broker (usually PHP class).
	Job string `json:"job"`

	// Ident is unique identifier of the job, should be provided from outside
	Ident string `json:"id"`

	// Payload is string data (usually JSON) passed to Job broker.
	Payload string `json:"payload"`

	// Headers with key-value pairs
	Headers map[string][]string `json:"headers"`

	// Options contains set of PipelineOptions specific to job execution. Can be empty.
	Options *Options `json:"options,omitempty"`
}

// Options carry information about how to handle given job.
type Options struct {
	// Priority is job priority, default - 10
	// pointer to distinguish 0 as a priority and nil as priority not set
	Priority int64 `json:"priority"`

	// Pipeline manually specified pipeline.
	Pipeline string `json:"pipeline,omitempty"`

	// Delay defines time duration to delay execution for. Defaults to none.
	Delay int64 `json:"delay,omitempty"`
}
