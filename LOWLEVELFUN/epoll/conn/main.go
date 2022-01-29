package main

import (
	"crypto/rand"
	"net"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		c, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}

		data := make([]byte, 70000)
		_, err = rand.Read(data)
		if err != nil {
			panic(err)
		}
		_, err = c.Write(data)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		//defer wg.Done()
		//c, err := net.Dial("tcp", "127.0.0.1:8080")
		//if err != nil {
		//	panic(err)
		//}
		//_, err = c.Write([]byte("wuzaaaa\n\r\n"))
		//if err != nil {
		//	panic(err)
		//}
	}()

	wg.Wait()
	time.Sleep(time.Second * 10)
}
