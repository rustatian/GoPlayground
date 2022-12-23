package main

import (
	"fmt"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"

	"github.com/rustatian/GoPlayground/temporal"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
		Logger:   NewZapAdapter(l),
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "", worker.Options{
		MaxConcurrentActivityExecutionSize: 3,
	})

	w.RegisterWorkflow(temporal.SampleTimerWorkflow)
	w.RegisterActivity(temporal.OrderProcessingActivity)
	//w.RegisterActivity(temporal.SendEmailActivity)
	//w.RegisterActivity(temporal.CurrentTime)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
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
