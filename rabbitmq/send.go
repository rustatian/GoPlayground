package main

import (
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
)

func FF(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	f, err := ioutil.ReadFile("/Users/0xdev/Projects/repo/GoPlayground/rabbitmq/file.txt")
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial("amqp://guest:guest@192.168.101.60:5672/")
	FF(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FF(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"bigfile", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, 	// routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        f,
		})
}
