package main

import (
	"context"
	"fmt"

	"github.com/centrifugal/centrifuge-go"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	client := centrifuge.NewProtobufClient("ws://localhost:8000/connection/websocket", centrifuge.Config{
		//GetToken: func(event centrifuge.ConnectionTokenEvent) (string, error) {
		//	println("toooken")
		//	token := subscriptionToken(event., "foo1", 0)
		//	return token, nil
		//},
		Token:              connToken("foo", 0),
		Data:               nil,
		Name:               "roadrunner_tests",
		Version:            "3.0.0",
		ReadTimeout:        0,
		WriteTimeout:       0,
		HandshakeTimeout:   0,
		MaxServerPingDelay: 0,
		TLSConfig:          nil,
		EnableCompression:  false,
		CookieJar:          nil,
		Header:             nil,
	})

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	client.OnSubscribing(func(event centrifuge.ServerSubscribingEvent) {
		println("sub")
	})

	sub, err := client.NewSubscription("bbbbb", centrifuge.SubscriptionConfig{
		Data: nil,
		//Token: connToken("foo", 0),
		GetToken: func(event centrifuge.SubscriptionTokenEvent) (string, error) {
			println("toooken")
			token := subscriptionToken(event.Channel, "foo", 0)
			return token, nil
		},
		Positioned:  false,
		Recoverable: true,
		JoinLeave:   true,
	})
	if err != nil {
		panic(err)
	}

	err = sub.Subscribe()
	if err != nil {
		panic(err)
	}

	pr, err := sub.Publish(context.Background(), []byte("foo"))
	if err != nil {
		panic(err)
	}

	fmt.Println(pr)
}

func connToken(user string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("e6ec932f-5319-4065-8014-302520abc5da"))
	if err != nil {
		panic(err)
	}
	return t
}

func subscriptionToken(channel string, user string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"channel": channel, "sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("e6ec932f-5319-4065-8014-302520abc5da"))
	if err != nil {
		panic(err)
	}
	return t
}

/*
  "connect_endpoint": "grpc://127.0.0.1:10001",
  "proxy_connect_endpoint": "grpc://127.0.0.1:10001",
  "proxy_connect_timeout":  "100s",

  "proxy_publish_endpoint": "grpc://127.0.0.1:10001",
  "proxy_publish_timeout":  "100s",

  "proxy_subscribe_endpoint": "grpc://127.0.0.1:10001",
  "proxy_subscribe_timeout":  "100s",

  "proxy_refresh_endpoint": "grpc://127.0.0.1:10001",
  "proxy_refresh_timeout": "100s",

  "proxy_rpc_endpoint": "grpc://127.0.0.1:10001",
  "proxy_rpc_timeout": "100s",
*/
