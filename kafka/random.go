package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close producer:", err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "test", Key: sarama.StringEncoder("key is set"), Value: sarama.StringEncoder("test")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to produce message to kafka cluster.")
	}

	log.Printf("Produced message to partition %d with offset %d", partition, offset)
}
