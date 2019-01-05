package main

import (
	"github.com/ValeryPiashchynski/go-2pc"
	"github.com/streadway/amqp"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

//type Coordinator interface {
//	Send(payload interface{}) error
//	Connect() error
//	Close() error
//}



func main() {
	a := &coordinator{}

	coord := go2pc.NewCoordinator(a)

	c, _ := sqlx.Connect("", "")
	tx, err := c.Begin()




	a := go2pc.NewGo2Pc()
}

type coordinator struct {
	ch *amqp.Channel
}

func(c *coordinator) Connect() error {

}

func(c *coordinator) Close() error {

}

func(c *coordinator) Send(payload interface{}) error {

}