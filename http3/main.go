package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/quic-go/quic-go/http3"

	"buf.build/gen/go/connectrpc/eliza/connectrpc/go/connectrpc/eliza/v1/elizav1connect"
	elizav1 "buf.build/gen/go/connectrpc/eliza/protocolbuffers/go/connectrpc/eliza/v1"
)

const (
	skipVerify = true
	// url        = "https://127.0.0.1:6660"

	// skipVerify = false
	url = "https://demo.connectrpc.com"
)

var reqBody = &elizav1.SayRequest{Sentence: "I feel happy."}

func main() {
	transport := &http3.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipVerify,
		},
	}
	defer transport.Close()
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Minute,
	}

	svcClient := elizav1connect.NewElizaServiceClient(client, url, connect.WithGRPC())

	log.Println("connect: ", url)
	log.Println("send: ", reqBody)
	resp, err := svcClient.Say(context.Background(), connect.NewRequest(reqBody))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Println("recv: ", resp.Msg)
}
