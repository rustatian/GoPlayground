package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Coalesce[T any](in chan T, maxBatch int, maxWait time.Duration, complete func([]T)) {
	items := make([]T, 0, maxBatch)
	for v := range in {
		items = append(items, v)

		t := time.NewTicker(maxWait)

	waitmore:
		for {
			select {
			case v, ok := <-in:
				if !ok {
					break waitmore
				}
				items = append(items, v)
				if len(items) >= maxBatch {
					break waitmore
				}
			case <-t.C:
				break waitmore
			}
		}
		t.Stop()

	fillbatch:
		for len(items) < maxBatch {
			select {
			case v, ok := <-in:
				if !ok {
					break fillbatch
				}
				items = append(items, v)
			default:
				break fillbatch
			}
		}

		complete(items)
		items = items[:0]
	}
}

func main() {
	q := make(chan int, 16)
	defer close(q)
	go Coalesce(q, 4, time.Second, func(batch []int) {
		fmt.Println(batch)
	})

	for i := 0; i < 1000; i++ {
		q <- i
		time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	}
}
