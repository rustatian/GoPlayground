package main

import (
	_ "embed"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	configImpl "github.com/roadrunner-server/config/v2"
	endure "github.com/roadrunner-server/endure/pkg/container"
	loggerImpl "github.com/roadrunner-server/logger/v2"
	serverImpl "github.com/roadrunner-server/server/v2"
)

//go:embed .rr.yaml
var rrYaml []byte

func main() {
	_ = os.Setenv("PATH", os.Getenv("PATH")+":"+os.Getenv("LAMBDA_TASK_ROOT"))
	_ = os.Setenv("LD_LIBRARY_PATH", "./lib:/lib64:/usr/lib64")

	cont, err := endure.NewContainer(nil, endure.SetLogLevel(endure.ErrorLevel))
	if err != nil {
		log.Fatal(err)
	}

	cfg := &configImpl.Plugin{}
	cfg.Timeout = time.Second * 30
	cfg.ReadInCfg = rrYaml
	cfg.Type = "yaml"

	// only 4 plugins needed here
	// 1. Server which should provide pool to us
	// 2. Our mini plugin, which expose this pool for us
	// 3. Logger
	// 4. Configurer
	err = cont.RegisterAll(
		cfg,
		&loggerImpl.Plugin{},
		&Plugin{},
		&serverImpl.Plugin{},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = cont.Init()
	if err != nil {
		log.Fatal(err)
	}

	ch, err := cont.Serve()
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case e := <-ch:
				log.Println(e.Error.Error())
			case <-sig:
				err = cont.Stop()
				if err != nil {
					log.Println(err)
				}
				return
			}
		}
	}()

	wg.Wait()
}
