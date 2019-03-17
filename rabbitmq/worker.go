package main

import (
	"bytes"
	"github.com/ValeryPiashchynski/GoPlayground/rabbitmq/ers"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func main() {

	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	ers.FailOnErrors(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()
	ers.FailOnErrors(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"hash",
		false,
		false,
		false,
		false,
		nil,
	)

	ers.FailOnErrors(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	ers.FailOnErrors(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
		}
	}()
	log.Printf("[*] Waiting for message")
	<-forever
}
