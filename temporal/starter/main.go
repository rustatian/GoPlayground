package main

import (
	"context"
	"log"

	"github.com/pborman/uuid"
	"github.com/rustatian/GoPlayground/temporal"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "greetings_" + uuid.New(),
		TaskQueue: "greetings",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, temporal.GreetingSample)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
