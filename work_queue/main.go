package main

import (
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"log"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

// Make an enqueuer with a particular namespace
var enqueuer = work.NewEnqueuer("my_app_namespace", redisPool)

func main() {
	// Enqueue a job named "send_email" with the specified parameters.
	_, err := enqueuer.Enqueue("send_email", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	if err != nil {
		log.Fatal(err)
	}
}
