package temporal

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// SampleTimerWorkflow workflow definition
func SampleTimerWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1000 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        0,
			BackoffCoefficient:     0,
			MaximumInterval:        0,
			MaximumAttempts:        100,
			NonRetryableErrorTypes: nil,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	workflow.ExecuteActivity(ctx, OrderProcessingActivity)
	//var data interface{}
	//err := f.Get(ctx, &data)
	//if err != nil {
	//	panic(err)
	//}

	//
	//queryType := "current_state"
	//err := workflow.SetQueryHandler(ctx, queryType, func() (string, error) {
	//	return time.Now().String(), nil
	//})
	//if err != nil {
	//	return err
	//}
	//
	////childCtx, cancelHandler := workflow.WithCancel(ctx)
	//selector := workflow.NewSelector(ctx)
	//
	//// In this sample case, we want to demo a use case where the workflow starts a long running order processing operation
	//// and in the case that the processing takes too long, we want to send out a notification email to user about the delay,
	//// but we won't cancel the operation. If the operation finishes before the timer fires, then we want to cancel the timer.
	//
	//var processingDone bool
	//f := workflow.ExecuteLocalActivity(ctx, OrderProcessingActivity)
	//selector.AddFuture(f, func(f workflow.Future) {
	//	processingDone = true
	//	// cancel timerFuture
	//	//cancelHandler()
	//})
	//
	//// use timer future to send notification email if processing takes too long
	//timerFuture := workflow.NewTimer(ctx, time.Second*60)
	//selector.AddFuture(timerFuture, func(f workflow.Future) {
	//	if !processingDone {
	//		// processing is not done yet when timer fires, send notification email
	//		_ = workflow.ExecuteLocalActivity(ctx, SendEmailActivity).Get(ctx, nil)
	//	}
	//})
	//
	//// wait the timer or the order processing to finish
	//selector.Select(ctx)
	//
	//// now either the order processing is finished, or timer is fired.
	//if !processingDone {
	//	// processing not done yet, so the handler for timer will send out notification email.
	//	// we still want the order processing to finish, so wait on it.
	//	selector.Select(ctx)
	//}

	workflow.GetLogger(ctx).Info("Workflow completed.")
	return nil
}

func OrderProcessingActivity(ctx context.Context) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("OrderProcessingActivity processing started.")
	//timeNeededToProcess := time.Second * 40
	//time.Sleep(timeNeededToProcess)
	//logger.Info("OrderProcessingActivity done.", "duration", timeNeededToProcess)

	time.Sleep(time.Second * 10)
	return "foo", nil
}

//func SendEmailActivity(ctx context.Context) error {
//	activity.GetLogger(ctx).Info("SendEmailActivity sending notification email as the process takes long time.")
//	return nil
//}
//
//func CurrentTime(_ context.Context) (string, error) {
//	return time.Now().String(), nil
//}
