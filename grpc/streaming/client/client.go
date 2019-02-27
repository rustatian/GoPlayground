package main

import (
	"context"
	"fmt"
	pb "github.com/ValeryPiashchynski/GoPlayground/grpc/streaming"
	"google.golang.org/grpc"
	"time"
)

//type fooClient struct {
//
//}
//
//func(c *fooClient) FooRPC(ctx context.Context, opts ...grpc.CallOption) (pb.FooService_FooRPCClient, error) {
//	grpc.NewClientStream(ctx,)
//}

//var serviceDescription = grpc.ServiceDesc{
//	ServiceName: "foo.FooService",
//	HandlerType: (*pb.FooServiceServer)(nil),
//	Methods:     []grpc.MethodDesc{},
//	Streams: []grpc.StreamDesc{
//		{
//			StreamName:    "FooRPC",
//			Handler:       pb.,
//			ServerStreams: true,
//			ClientStreams: true,
//		},
//	},
//	Metadata: "foo.proto",
//}

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:30000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	client := pb.NewFooServiceClient(conn)
	stream, err := client.FooRPC(ctx)

	wc := make(chan struct{})

	msg := &pb.Data{Msg:"some_data"}
	go func() {
		for {
			fmt.Println("Sleep for 1 second")
			time.Sleep(time.Second * 1)
			fmt.Println("Sending message")
			err := stream.SendMsg(msg)
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