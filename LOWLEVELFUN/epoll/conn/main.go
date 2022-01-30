package main

import (
	"crypto/rand"
	"fmt"
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

		data := make([]byte, 7000)
		_, err = rand.Read(data)
		if err != nil {
			panic(err)
		}
		_, err = c.Write(data)
		if err != nil {
			panic(err)
		}

		res := make([]byte, 7000)
		c.Read(res)
		fmt.Printf("res: %s\n", res)
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
