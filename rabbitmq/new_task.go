package main

import (
	"github.com/ValeryPiashchynski/GoPlayground/rabbitmq/ers"
	"github.com/streadway/amqp"
	"os"
	"strings"
)

func main() {
	body := bodyFrom(os.Args)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ers.FailOnErrors(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	ers.FailOnErrors(err, "Failed to open a channel")
	defer ch.Close()

	//q, err := ch.QueueDeclare(
	//	"hello",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//ers.FailOnErrors(err, "Failed to declare a queue")

	err = ch.Publish(
		"",
		"hash",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
