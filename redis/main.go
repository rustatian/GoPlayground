package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func main() {
	pool := newRedisPool("localhost:6379", 10, "valery")

	conn := pool.Get()
	reply, err := conn.Do("SET", "key", "whatever")
	if err != nil {
		panic(err)
	}

	fmt.Print(reply)
}

func newRedisPool(address string, threads int, namespace string) *redis.Pool {
	return &redis.Pool{
		MaxActive:   threads,
		MaxIdle:     10,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
		Wait: true,
	}
}

//reply, err := conn.Do("SET", b)
