package main

import (
	"log"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/rustatian/GoPlayground/temporal"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "default", worker.Options{
		MaxConcurrentActivityExecutionSize: 3,
	})

	w.RegisterWorkflow(temporal.SampleTimerWorkflow)
	w.RegisterActivity(temporal.OrderProcessingActivity)
	w.RegisterActivity(temporal.SendEmailActivity)
	w.RegisterActivity(temporal.CurrentTime)

	ch := make(chan interface{})

	go func() {
		time.Sleep(time.Second * 50)
		ch <- struct{}{}
	}()

	err = w.Run(ch)
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
