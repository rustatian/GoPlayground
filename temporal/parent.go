package temporal

import (
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/workflow"
)

func SampleParentWorkflow(ctx workflow.Context) (string, error) {
	logger := workflow.GetLogger(ctx)

	cwo := workflow.ChildWorkflowOptions{
		ParentClosePolicy: enumspb.PARENT_CLOSE_POLICY_TERMINATE,
		WorkflowID:        "ABC-SIMPLE-CHILD-WORKFLOW-ID",
	}
	ctx = workflow.WithChildOptions(ctx, cwo)

	var result string
	err := workflow.ExecuteChildWorkflow(ctx, SampleChildWorkflow, "World").Get(ctx, &result)
	if err != nil {
		logger.Error("Parent execution received child execution failure.", "Error", err)
		return "", err
	}

	logger.Info("Parent execution completed.", "Result", result)
	return result, nil
}
