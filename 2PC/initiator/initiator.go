package main

import (
	"github.com/ValeryPiashchynski/go-2pc"
	"github.com/ValeryPiashchynski/go-2pc/transport"
	_ "github.com/lib/pq"
	"log"
	"os"
)

//type Coordinator interface {
//	Send(payload interface{}) error
//	Connect() error
//	Close() error
//}

//var log = logrus.New()

func main() {
	//log.Out = os.Stdout

	l := log.New(os.Stdout, "", 0)
	a := transport.NewAMQPCoordinator("amqp://guest:guest@localhost:5672/", 4)

	b := go2pc.NewInitiator(a, nil, l,)

	err := b.InitCoordinator()
	if err != nil {
		panic(err)
	}

	err = b.Begin2PCCommit()
	if err != nil {
		panic(err)
	}
}
