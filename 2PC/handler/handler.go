package main

import (
	"crypto/rand"
	"fmt"
	"github.com/ValeryPiashchynski/go-2pc"
	"github.com/ValeryPiashchynski/go-2pc/transport/amqp"
	"io"
	"log"
	"os"
)

func main() {
	l := log.New(os.Stdout, "", 0)
	a := amqp.NewAMQPCoordinator("amqp://guest:guest@localhost:5672/", )

	h := go2pc.New2PCHandler(a, l)

	err := h.InitCoordinator()
	if err != nil {
		panic(err)
	}

	// signal -
	// replyTo
	err = h.PhaseDone(go2pc.InitDone, go2pc.Init2PC, "temp1")
	if err != nil {
		panic(err)
	}

	//phase 1

	err = h.PhaseDone(go2pc.Phase1Done, go2pc.Phase1Start, "temp2")
	if err != nil {
		panic(err)
	}

	// tx.commit

	err = h.PhaseDone(go2pc.Phase2Done, go2pc.Phase2Start, "temp3")
	if err != nil {
		panic(err)
	}
}

func randomBoundary() string {
	var buf [30]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", buf[:])
}
