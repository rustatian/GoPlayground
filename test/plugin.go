package main

import (
	"context"
	"sync"
	"unsafe"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/spiral/roadrunner/v2/pkg/payload"
	"github.com/spiral/roadrunner/v2/pkg/pool"
	"github.com/spiral/roadrunner/v2/plugins/logger"
	"github.com/spiral/roadrunner/v2/plugins/server"
)

type Plugin struct {
	sync.Mutex
	log     logger.Logger
	srv     server.Server
	wrkPool pool.Pool
}

func (p *Plugin) Init(srv server.Server, log logger.Logger) error {
	var err error
	p.srv = srv
	p.log = log
	return err
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)
	p.Lock()
	defer p.Unlock()
	var err error

	p.wrkPool, err = p.srv.NewWorkerPool(context.Background(), pool.Config{
		Debug:           false,
		NumWorkers:      1,
		MaxJobs:         0,
		AllocateTimeout: 0,
		DestroyTimeout:  0,
		Supervisor: &pool.SupervisorConfig{
			WatchTick:       0,
			TTL:             0,
			IdleTTL:         0,
			ExecTTL:         0,
			MaxWorkerMemory: 0,
		},
	}, nil, nil)

	go func() {
		// register handler
		lambda.Start(p.handler())
	}()

	if err != nil {
		errCh <- err
	}
	return errCh
}

func (p *Plugin) Stop() error {
	p.Lock()
	defer p.Unlock()

	if p.wrkPool != nil {
		p.wrkPool.Destroy(context.Background())
	}
	return nil
}

func (p *Plugin) handler() func(pld string) (string, error) {
	return func(pld string) (string, error) {
		data := fastConvert(pld)
		// execute on worker pool
		if p.wrkPool == nil {
			// or any error
			return "", nil
		}
		exec, err := p.wrkPool.Exec(payload.Payload{
			Context: nil,
			Body:    data,
		})
		if err != nil {
			return "", err
		}
		return exec.String(), nil
	}
}

// reinterpret_cast conversion cast from string to []byte
// unsafe
func fastConvert(d string) []byte {
	return *(*[]byte)(unsafe.Pointer(&d))
}
