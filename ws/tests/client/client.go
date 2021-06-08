package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	_ "embed"

	"github.com/48d90782/GoPlayground/ws/tests/message"
	"github.com/fasthttp/websocket"
)

var addr = flag.String("addr", "localhost:15395", "http service address")
var redial = flag.Bool("r", true, "reconnect when lost connection")
var numOfClients = flag.Uint64("n", 1000, "number of connections")
var scheme = flag.String("s", "ws", "websocket scheme")

var numberOfActiveClients int64 = 0
var totalNumberOfMessages uint64 = 0
var stop uint64 = 0

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	flag.Parse()

	for i := uint64(0); i < *numOfClients; i++ {
		go func() {
			conn := NewWSConnection()
			conn.dial()
		}()
	}

	go func() {
		tt := time.NewTicker(time.Second)
		for {
			select {
			case <-tt.C:
				fmt.Printf("Actual number of the WS connections: %d, rate: %d msg/s\n", atomic.LoadInt64(&numberOfActiveClients), atomic.LoadUint64(&totalNumberOfMessages))
				// clear
				atomic.StoreUint64(&totalNumberOfMessages, 0)
			}
		}
	}()

	<-interrupt
	*redial = false
	fmt.Println("------------------------ received stop signal ------------------------")
	atomic.AddUint64(&stop, 1)

	// exit when all connections will be closed
	for {
		if atomic.CompareAndSwapInt64(&numberOfActiveClients, 0, 0) {
			fmt.Println("---------- ALL CLIENTS DISCONNECTED --------------")
			return
		}
	}
}

type Connection struct {
	url url.URL

	redialCh chan struct{}

	// internal
	conn net.Conn
}

// NewWSConnection data to dial should be provide before
func NewWSConnection() *Connection {
	conn := &Connection{
		url:      url.URL{Scheme: *scheme, Host: *addr, Path: "/ws"},
		redialCh: make(chan struct{}, 4),
	}
	conn.redialHandler()
	return conn
}

func (conn *Connection) redialHandler() {
	go func() {
		for {
			select {
			case <-conn.redialCh:
				if !*redial {
					return
				}
				// redial
				println("------------------ GOT REDIAL SIGNAL -------------------------")
				go conn.dial()
			}
		}
	}()
}

func (conn *Connection) dial() {
	log.Printf("connecting to %s", conn.url.String())

	da := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: time.Second * 20,
	}

	stopCh := make(chan struct{}, 2)
	stCh := make(chan struct{}, 2)

	c, resp, err := da.Dial(conn.url.String(), nil)
	if err != nil {
		log.Println(err)
		conn.redialCh <- struct{}{}
		return
	}

	defer func() {
		_ = c.Close()
		_ = resp.Body.Close()
	}()

	go func() {
		for {
			select {
			case <-stCh:
				return
			default:
				time.Sleep(time.Second * 5)
				if atomic.CompareAndSwapUint64(&stop, 1, 1) {
					stopCh <- struct{}{}
					return
				}
			}
		}
	}()

	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				stopCh <- struct{}{}
				return
			}

			atomic.AddUint64(&totalNumberOfMessages, 1)
			_, _ = io.Discard.Write(msg)
		}
	}()

	data := "hello bbbbbbbbbbbbbbbeeeeeeeeeee"
	m := &message.Message{
		Topics:  []string{"foo", "foo2"},
		Command: "join",
		Broker:  "memory",
		Payload: []byte(data),
	}

	d, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	// Cleanly close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	err = c.WriteMessage(websocket.BinaryMessage, d)
	if err != nil {
		log.Println("error update", err)
		conn.redialCh <- struct{}{}
		return
	}

	// add active client
	atomic.AddInt64(&numberOfActiveClients, 1)
	// remove client when disconnected
	defer atomic.AddInt64(&numberOfActiveClients, ^int64(0))

	// wait for stop
	<-stopCh
	// stop atomicCSW goroutine
	stCh <- struct{}{}

	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
	}

	conn.redialCh <- struct{}{}
}
