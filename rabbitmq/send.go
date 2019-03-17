package main

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func FF(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FF(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FF(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"refine_inventory", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		panic(err)
	}

	tt := UpdateRequest{
		CompanyID: "3333324234234213412341234",
		Type:      "company",
		Scope:     InventoryScope{[]string{"1"}, []string{"1"}, []string{"1"}}}
	b, _ := tt.Marshal()

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Headers:     map[string]interface{}{"JWTToken": "Bearer ",},
			ContentType: "text/plain",
			Body:        b,
		})
}

type InventoryScope struct {
	GroupIds   []string
	PackageIds []string
	ProductIds []string
}

//UpdateRequest is used as notify message to the productservice
type UpdateRequest struct {
	CompanyID string         `json:"companyId"`
	Type      UpdateType     `json:"type"`
	Scope     InventoryScope `json:"scope"`
}

//Marshal is using for hiding functionality of converting struct to []byte json
func (u *UpdateRequest) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

//UpdateType is used to separate company and other update requests
type UpdateType string
