package main

import (
	"context"
	"fmt"
	pb "github.com/ValeryPiashchynski/GoPlayground/grpc/streaming"
	"google.golang.org/grpc"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*500)

	conn, err := grpc.Dial("localhost:30000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	client := pb.NewFooServiceClient(conn)
	stream, err := client.FooRPC(ctx)

	wc := make(chan struct{})

	buf := make([]string, 0)
	go func() {
		for {
			select {
			case <-ctx.Done():
				panic(ctx)
			default:
				time.Sleep(time.Millisecond * 100)
				s := "flasfjash;fjhas;ljdf;lasjdf;ljasdf;ljlfaj;lsaghwret235dfsaddfj;"
				j := ""
				for i := 0; i < 1000; i++ {
					j += s
				}
				msg := &pb.Data{Msg: j}
				err := stream.SendMsg(msg)
				buf = append(buf, j)
				if err != nil {
					cancel()
					fmt.Println(err)
				}
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
