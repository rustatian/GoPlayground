package main

import (
	"github.com/ValeryPiashchynski/go-2pc"
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
		"2PC_Work", // queue
		"",    // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	someService := "someService"
	//someError := "some very very complicated error"
	r := go2pc.Response{
		ServiceName: &someService,
		//Err:         &someError,
	}
	b, _ := r.Marshal()

	forever := make(chan bool)

	go func() {
		for l := range msgs {
			switch l.CorrelationId {
			// init phase
			case go2pc.Init2PC:
				log.Printf("Received a message: %s", l.Body)
				log.Printf("Received a message exchange is: %s", l.Exchange)

				err = ch.Publish(
					"2PHASE_COMMIT", // exchange
					"2PC_Work", // routing key
					false, // mandatory
					false, // immediate
					amqp.Publishing{
						ContentType:   "application/json",
						CorrelationId: go2pc.InitDone,
						Body:          b,
					})

				err = l.Ack(false)
				if err != nil {
					panic(err)
				}
				return
			}
		}
	}()

	//go listen(msgs)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
