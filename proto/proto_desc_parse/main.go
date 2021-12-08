package main

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func main() {
	desc, err := protoregistry.GlobalFiles.FindFileByPath("test.proto")
	if err != nil {
		panic(err)
	}

	// err = protoregistry.GlobalFiles.RegisterFile(desc)
	// if err != nil {
	// 	panic(err)
	// }

	meth := desc.Options().ProtoReflect().ProtoMethods()
	_ = meth

	m := desc.Services()
	nm := m.Get(0)

	// serviceDesc := &grpc.ServiceDesc{
	// 	Methods: nm.Name(),
	// }

	println(nm)

	server := grpc.NewServer()
	server.RegisterService(nil, nil)

	// methods := desc.Options().ProtoReflect().ProtoMethods()

	// server.RegisterService(desc.Services().ByName(nil), nil)

	fmt.Println(desc)
}

func Load(svc *grpc.ServiceDesc) {

}
