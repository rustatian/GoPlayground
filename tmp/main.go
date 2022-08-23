package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	rrLib "github.com/roadrunner-server/roadrunner/v2/lib"
)

// go get -u github.com/roadrunner-server/roadrunner/v2
func main() {
	rr, err := rrLib.NewRR("/path/to/.rr.yaml", nil, rrLib.DefaultPluginsList())
	if err != nil {
		panic(err)
	}

	stopCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	go func() {
		err2 := rr.Serve()
		if err2 != nil {
			errCh <- err2
		}
	}()

	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case e := <-errCh:
		fmt.Printf("error occured: %v\n", e)
		return
	case <-stopCh:
		rr.Stop()
	}
}
