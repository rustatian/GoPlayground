package main

import (
	"github.com/streadway/amqp"
	"log"
)

func FF(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.101.60:5672/")
	FF(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FF(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"a",   // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	//mm := amqp.Publishing{
	//	ContentType:  "text/plain",
	//	Body:         []byte("Go Go AMQP!"),
	//}
	//
	//err = ch.Publish("", q.Name, false, true, mm)
	//if err != nil {
	//	panic(err)
	//}

	FF(err, "Failed to declare a queue")

	//q2, err := ch.QueueDeclare(
	//	"b", // name
	//	false,    // durable
	//	false,    // delete when unused
	//	false,    // exclusive
	//	false,    // no-wait
	//	nil,      // arguments
	//)

	body := "Hello World!"
	//body2 := "Hello World2!"
	err = ch.Publish(
		"INTURN", // exchange
		q.Name,   // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		panic(err)
	}

	//err = ch.Publish(
	//	"",      // exchange
	//	q2.Name, // routing key
	//	false,   // mandatory
	//	false,   // immediate
	//	amqp.Publishing{
	//		ContentType: "text/plain",
	//		Body:        []byte(body2),
	//	})
	//log.Printf(" [x] Sent %s", body)
	//FF(err, "Failed to publish a message")
}
