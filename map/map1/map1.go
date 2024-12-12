package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	sm := sync.Map{}

	for i := 0; i < 100000000; i++ {
		data := &struct {
			Key   string
			Value []byte
		}{
			Key:   "a",
			Value: []byte("aa"),
		}
		sm.Store("foo", data)
		go func() {
			time.Sleep(time.Millisecond * 100)
			sm.Delete("foo")
		}()

		fmt.Printf("num goroutines: %d\n", runtime.NumGoroutine())
	}

	time.Sleep(time.Second * 30)
}
