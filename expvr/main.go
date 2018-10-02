package main

import (
	_ "expvar"
	"fmt"
	"net"
	"net/http"
)

func main() {
	sock, err := net.Listen("tcp", "localhost:8123")
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP now available at port 8123")
	http.Serve(sock, nil)
}
