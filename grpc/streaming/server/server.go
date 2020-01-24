package main

import (
	"fmt"
	pb "github.com/ValeryPiashchynski/GoPlayground/grpc/streaming"
	"google.golang.org/grpc"
	"io"
	"net"
)

type fooServer struct {

}

func(f *fooServer) FooRPC(stream pb.FooService_FooRPCServer) error {

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Message: " + msg.Msg[:100])
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterFooServiceServer(server,&fooServer{})
	listener, err := net.Listen("tcp", ":30000")
	if err != nil {
		fmt.Println(err)
	}
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
	}
}