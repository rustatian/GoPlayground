package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	msgs, err := ch.Consume(
		"abort", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for l := range msgs {
			switch l.CorrelationId {
			// init phase
			case "init_2PC":
				log.Printf("Received a message: %s", l.Body)
				log.Printf("Received a message Type is: %s", l.Type)
				l.Ack(false)

			}

		}
	}()


	//go listen(msgs)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func listen(delivery <-chan amqp.Delivery) {
	for {
		select {
		case m, _ := <-delivery:
			if m.CorrelationId == "123" {
				log.Printf("Received a message: %s", m.Body)
				log.Printf("Received a message Type is: %s", m.Type)
			}
		}
	}
}