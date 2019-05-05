package main

import (
	"github.com/pkg/errors"
	"time"
)

func main() {
	time.Sleep(time.Second * 5)

	panic(errors.New("Whooops, error!"))
}
