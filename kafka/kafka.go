package main

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	config := sarama.NewConfig()
	//config.Producer.MaxMessageBytes = 100000000
	//config.Producer.Return.Successes = true
	//config.ChannelBufferSize = 1000000000
	//config.Net.MaxOpenRequests = 100000
	producer, err := sarama.NewAsyncProducer([]string{"54.152.193.114:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for a := range producer.Successes() {
			aa, _ := a.Value.Encode()
			log.Println(string(aa))
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range producer.Errors() {
			log.Println(i.Err)
			errors++
		}
	}()

	for {
		message := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}
