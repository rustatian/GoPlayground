package temporal

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// @@@SNIPSTART samples-go-child-workflow-example-child-workflow-definition
// SampleChildWorkflow is a Workflow Definition
func SampleChildWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)

	time.Sleep(time.Second * 30)

	greeting := "Hello " + name + "!"
	logger.Info("Child workflow execution: " + greeting)
	return greeting, nil
}
