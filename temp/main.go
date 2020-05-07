package main

import "net"

func main() {
	_, err := net.Dial("unix", "/home/valery/Projects/opensource/goridge/tests/server.sock")
	if err != nil {
		panic(err)
	}
}
