package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	a := amqp.Table{}
	a = nil
	fmt.Print(fmt.Sprintf("%v", a))
	//queue := make(chan string)
	//go func() {
	//
	//	queue <- "one"
	//	queue <- "two"
	//	close(queue)
	//}()

	//for elem := range queue {
	//	fmt.Println(elem)
	//}
}
