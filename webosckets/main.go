package main

import (
	"bufio"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/ws", func(writer http.ResponseWriter, request *http.Request) {
	})
}

type Packet struct {
}

type Channel struct {
	conn net.Conn
	send chan Packet
}

func NewChannel(conn net.Conn) *Channel {
	c := &Channel{
		conn: conn,
		send: make(chan Packet, 100),
	}

	go c.reader()
	go c.writer()
	return c
}

func (c *Channel) reader() {
	buf := bufio.NewReader(c.conn)
	for {
		pkt, _ := readPacket(buf)
		c.handle(pkg)
	}
}

func (c *Channel) writer() {
	buf := bufio.NewWriter(c.conn)

	for pkt := range c.send {
		_ = writePacket(pkt)
		_ = buf.Flush()
	}
}
