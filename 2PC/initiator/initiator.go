package main

import (
	"github.com/ValeryPiashchynski/go-2pc"
	"github.com/ValeryPiashchynski/go-2pc/transport/amqp"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	//log.Out = os.Stdout

	l := log.New(os.Stdout, "", 0)
	a := amqp.NewAMQPCoordinator("amqp://guest:guest@localhost:5672/", )

	b := go2pc.NewInitiator(a, 1,nil, l,)

	err := b.InitCoordinator()
	if err != nil {
		panic(err)
	}

	err = b.Prepare(go2pc.Init2PC, go2pc.InitDone,)
	if err != nil {
		panic(err)
	}
	err = b.StartPhase(go2pc.Phase1Start, go2pc.Phase1Done,)
	if err != nil {
		panic(err)
	}
	err = b.StartPhase(go2pc.Phase2Start, go2pc.Phase2Done,)
	if err != nil {
		panic(err)
	}
}
