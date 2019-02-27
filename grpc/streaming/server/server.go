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
	fmt.Println("Start a stream")

	for {
		msg, err := stream.Recv()
		fmt.Println("Recieved a message")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("Message: " + msg.Msg)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterFooServiceServer(server,&fooServer{})
	listener, err := net.Listen("tcp", ":30000")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}