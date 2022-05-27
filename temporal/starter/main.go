package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	"github.com/rustatian/GoPlayground/temporal"
	"go.temporal.io/sdk/client"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "timer_" + uuid.New(),
		TaskQueue: "default",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, temporal.SampleTimerWorkflow)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
