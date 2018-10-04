package main

import "GoPlayground/rabbitmq/rabbitMQ"

func main() {
	rabbit.NewRabbit("", "", "", true, true, true, true, nil).
		Connect().
		CreateChannel().
		Handler( /*gokit.handler with enpoint/decode/encode*/ ).
		StartListen()
	if err != nil {

	}
	print(msg)

}
