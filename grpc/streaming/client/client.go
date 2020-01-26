package main

import (
	"context"
	"fmt"
	pb "github.com/ValeryPiashchynski/GoPlayground/grpc/streaming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net/http"
	_ "net/http/pprof"
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
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6061", nil))
	}()

	conn, err := grpc.Dial("localhost:30000", grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                12 * time.Second,
		Timeout:             12 * time.Second,
		PermitWithoutStream: false,
	}))
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 30)
	client := pb.NewFooServiceClient(conn)
	stream, err := client.FooRPC(context.Background())

	wc := make(chan struct{})

	buf := make([]string, 0)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			s := "flasfjash;fjhas;ljdf;lasjdf;ljasdf;ljlfaj;lsaghwret235dfsaddfj;"
			j := ""
			for i := 0; i < 1000; i++ {
				j += s
			}
			msg := &pb.Data{Msg: j}
			err := stream.SendMsg(msg)
			buf = append(buf, j)
			if err != nil {
				wc <- struct{}{}
				fmt.Println(err)
			}

		}
	}()
	<-wc
	err = stream.CloseSend()
	if err != nil {
		fmt.Println(err)
	}

	err = conn.Close()
	if err != nil {
		fmt.Println(err)
	}
}
