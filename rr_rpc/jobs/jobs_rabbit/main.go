package main

import (
	"flag"
	"log"

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

	msg, err := pack("212", &Item{
		Job:     "super-job",
		Ident:   "212",
		Payload: "asdfkasjdf;laskjdflkajsdfl;kjas;ldkfjlksjdlfkasjdf;lkajs;dkjfkl;asjdfl",
		Headers: nil,
		Options: &Options{
			Priority:   1,
			Pipeline:   "test-1",
			Delay:      0,
			Attempts:   0,
			RetryDelay: 0,
			Timeout:    0,
		},
		AckFunc:  nil,
		NackFunc: nil,
	})

	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan struct{})

	for j := 0; j < 100; j++ {
		go func() {
			for i := 0; i < 1_000_000; i++ {
				err = channel.Publish(
					"default",
					"test",
					false,
					false,
					amqp.Publishing{
						Headers:     msg,
						ContentType: "application/octet-stream",
						Body:        []byte("asdfal;sjf;lasjdf;lkjaal;sdfj;lkajsdl;fk"),
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

const (
	rrID         string = "rr_id"
	rrJob        string = "rr_job"
	rrHeaders    string = "rr_headers"
	rrPipeline   string = "rr_pipeline"
	rrTimeout    string = "rr_timeout"
	rrDelay      string = "rr_delay"
	rrRetryDelay string = "rr_retry_delay"
)

// pack job metadata into headers
func pack(id string, j *Item) (amqp.Table, error) {
	headers, err := json.Marshal(j.Headers)
	if err != nil {
		return nil, err
	}
	return amqp.Table{
		rrID:         id,
		rrJob:        j.Job,
		rrPipeline:   j.Options.Pipeline,
		rrHeaders:    headers,
		rrTimeout:    j.Options.Timeout,
		rrDelay:      j.Options.Delay,
		rrRetryDelay: j.Options.RetryDelay,
	}, nil
}

type Item struct {
	// Job contains pluginName of job broker (usually PHP class).
	Job string `json:"job"`

	// Ident is unique identifier of the job, should be provided from outside
	Ident string

	// Payload is string data (usually JSON) passed to Job broker.
	Payload string `json:"payload"`

	// Headers with key-values pairs
	Headers map[string][]string

	// Options contains set of PipelineOptions specific to job execution. Can be empty.
	Options *Options `json:"options,omitempty"`

	// Ack delegates an acknowledgement through the Acknowledger interface that the client or server has finished work on a delivery
	AckFunc func(multiply bool) error

	// Nack negatively acknowledge the delivery of message(s) identified by the delivery tag from either the client or server.
	// When multiple is true, nack messages up to and including delivered messages up until the delivery tag delivered on the same channel.
	// When requeue is true, request the server to deliver this message to a different consumer. If it is not possible or requeue is false, the message will be dropped or delivered to a server configured dead-letter queue.
	// This method must not be used to select or requeue messages the client wishes not to handle, rather it is to inform the server that the client is incapable of handling this message at this time
	NackFunc func(multiply bool, requeue bool) error
}

// Options carry information about how to handle given job.
type Options struct {
	// Priority is job priority, default - 10
	// pointer to distinguish 0 as a priority and nil as priority not set
	Priority uint32 `json:"priority"`

	// Pipeline manually specified pipeline.
	Pipeline string `json:"pipeline,omitempty"`

	// Delay defines time duration to delay execution for. Defaults to none.
	Delay int32 `json:"delay,omitempty"`

	// Attempts define maximum job retries. Attention, value 1 will only allow job to execute once (without retry).
	// Minimum valuable value is 2.
	Attempts int32 `json:"maxAttempts,omitempty"`

	// RetryDelay defines for how long job should be waiting until next retry. Defaults to none.
	RetryDelay int32 `json:"retryDelay,omitempty"`

	// Reserve defines for how broker should wait until treating job are failed. Defaults to 30 min.
	Timeout int32 `json:"timeout,omitempty"`
}
