package main

import (
	"github.com/ValeryPiashchynski/go-2pc"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	someService := "someService"
	//someError := "some very very complicated error"
	r := go2pc.Response{
		ServiceName: &someService,
		//Err:         &someError,
	}
	b, _ := r.Marshal()
	err = ch.Publish(
		"2PHASE_COMMIT",    // exchange
		"2PC_Work", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: go2pc.InitDone,
			Body:          b,
		})
}
