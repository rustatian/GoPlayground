package main

import (
	"log"

	"github.com/streadway/amqp"
)

func F(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.101.60:5672/")
	F(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	F(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"a",   // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	//q2, err := ch.QueueDeclare(
	//	"b", // name
	//	false,    // durable
	//	false,    // delete when unused
	//	false,    // exclusive
	//	false,    // no-wait
	//	nil,      // arguments
	//)
	F(err, "Failed to declare a queue")

	ch.QueueBind(q.Name, "", "INTURN", false, nil)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	F(err, "Failed to register a consumer")

	//msgs2, err := ch.Consume(
	//	q2.Name, // queue
	//	"",      // consumer
	//	true,    // auto-ack
	//	false,   // exclusive
	//	false,   // no-local
	//	false,   // no-wait
	//	nil,     // args
	//)

	forever := make(chan bool)

	//go func() {
	//	for {
	//		select {
	//		case m, _ := <-msgs:
	//			log.Printf("Received a message: %s", m.Body)
	//		default:
	//
	//		}
	//	}
	//}()

	go listen(msgs)

	//go func() {
	//	for d := range msgs2 {
	//		log.Printf("Received a message: %s", d.Body)
	//		return
	//	}
	//}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func listen(delivery <-chan amqp.Delivery) {
	for {
		select {
		case m, _ := <-delivery:
			go func(msg amqp.Delivery) {
				log.Printf("Received a message: %s", m.Body)
				log.Printf("Received a message Type is: %s", m.Type)
			}(m)
		default:

		}
	}
}
