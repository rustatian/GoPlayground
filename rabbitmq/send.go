package main

import (
	"github.com/ValeryPiashchynski/GoPlayground/rabbitmq/ers"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	ers.FailOnErrors(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	ers.FailOnErrors(err, "Failed to open a channel")
	defer ch.Close()

	//q, err := ch.QueueDeclare(
	//	"hash",
	//	false,
	//	false,
	//	false,
	//	false,
	//	nil,
	//)
	//ers.FailOnErrors (err, "Failed to declare a queue")

	body := `{"password":"20574178"}`
	err = ch.Publish(
		"",
		"hash",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
}
