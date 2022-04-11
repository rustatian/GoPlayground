package main

import (
	"log"

	"github.com/rustatian/GoPlayground/temporal"
	"github.com/uber-go/tally/v4/prometheus"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	ms, cl, errPs := temporal.NewPrometheusScope(prometheus.Configuration{
		ListenAddress: "localhost:9095",
		TimerType:     "summary",
	}, "roadrunner_", l)
	if errPs != nil {
		panic(errPs)
	}

	defer func() {
		_ = cl.Close()
	}()

	mh := tally.NewMetricsHandler(ms)
	c, err := client.NewClient(client.Options{
		HostPort:       client.DefaultHostPort,
		MetricsHandler: mh,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greetings", worker.Options{})

	w.RegisterWorkflow(temporal.GreetingSample)
	activities := &temporal.Activities{Name: "Temporal", Greeting: "Hello"}
	w.RegisterActivity(activities)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
