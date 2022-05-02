package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/rustatian/GoPlayground/temporal"
)

// @@@SNIPSTART samples-go-child-workflow-example-worker-starter
func main() {
	// The client is a heavyweight object that should be created only once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "child-workflow", worker.Options{})

	w.RegisterWorkflow(temporal.SampleParentWorkflow)
	w.RegisterWorkflow(temporal.SampleChildWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
