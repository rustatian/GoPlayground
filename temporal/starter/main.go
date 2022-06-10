package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pborman/uuid"
	"github.com/rustatian/GoPlayground/temporal"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
		Logger:   NewZapAdapter(l),
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

	var in any
	err = we.Get(context.Background(), &in)
	if err != nil {
		panic(err)
	}
}

type ZapAdapter struct {
	zl *zap.Logger
}

// NewZapAdapter ... which uses general log interface
func NewZapAdapter(zapLogger *zap.Logger) *ZapAdapter {
	return &ZapAdapter{
		zl: zapLogger.WithOptions(zap.AddCallerSkip(1)),
	}
}

func (log *ZapAdapter) Debug(msg string, keyvals ...interface{}) {
	log.zl.Debug(msg, log.fields(keyvals)...)
}

func (log *ZapAdapter) Info(msg string, keyvals ...interface{}) {
	log.zl.Info(msg, log.fields(keyvals)...)
}

func (log *ZapAdapter) Warn(msg string, keyvals ...interface{}) {
	log.zl.Warn(msg, log.fields(keyvals)...)
}

func (log *ZapAdapter) Error(msg string, keyvals ...interface{}) {
	log.zl.Error(msg, log.fields(keyvals)...)
}

func (log *ZapAdapter) fields(keyvals []interface{}) []zap.Field {
	// we should have even number of keys and values
	if len(keyvals)%2 != 0 {
		return []zap.Field{zap.Error(fmt.Errorf("odd number of keyvals pairs: %v", keyvals))}
	}

	zf := make([]zap.Field, len(keyvals)/2)
	j := 0
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			key = fmt.Sprintf("%v", keyvals[i])
		}

		zf[j] = zap.Any(key, keyvals[i+1])
		j++
	}

	return zf
}
