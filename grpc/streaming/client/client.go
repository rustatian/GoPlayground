package client

import (
	"context"
	"fmt"
	pb "github.com/ValeryPiashchynski/GoPlayground/grpc/streaming"
	"google.golang.org/grpc"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:30000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	client := pb.NewFooServiceClient(conn)
	stream, err := client.FooRPC(ctx)

	wc := make(chan struct{})

	msg := &pb.Data{Msg:"some_data"}
	go func() {
		for {
			fmt.Println("Sleep for 1 second")
			time.Sleep(time.Second * 1)
			fmt.Println("Sending message")
			err := stream.Send(msg)
			if err != nil {
				panic(err)
			}
		}
	}()

	<- wc
	err = stream.CloseSend()
	if err != nil {
		panic(err)
	}
}