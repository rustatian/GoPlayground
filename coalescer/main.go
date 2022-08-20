package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Coalescer[T any] struct {
	state    chan state[T]
	maxBatch int
	maxWait  time.Duration
	complete func([]T)
}

type state[T any] struct {
	items          []T
	firstItemAdded time.Time
	timer          *time.Timer
}

func NewCoalescer[T any](maxBatch int, maxWait time.Duration, complete func([]T)) *Coalescer[T] {
	c := make(chan state[T], 1)
	c <- state[T]{}
	return &Coalescer[T]{c, maxBatch, maxWait, complete}
}

func (c *Coalescer[T]) Put(item T) {
	s := <-c.state
	defer func() { c.state <- s }()

	s.items = append(s.items, item)
	if len(s.items) >= c.maxBatch {
		s.flush(c.complete)
		return
	}

	if len(s.items) > 1 {
		return
	}
	s.firstItemAdded = time.Now()

	if s.timer != nil {
		s.timer.Reset(c.maxWait)
		return
	}
	s.timer = time.AfterFunc(c.maxWait, func() {
		s := <-c.state
		if len(s.items) > 0 {
			elapsed := time.Since(s.firstItemAdded)
			if elapsed >= c.maxWait {
				s.flush(c.complete)
			} else {
				s.timer.Reset(c.maxWait - elapsed)
			}
		}
		c.state <- s
	})
}

func (c *Coalescer[T]) Flush() {
	s := <-c.state
	if len(s.items) > 0 {
		s.flush(c.complete)
	}
	c.state <- s
}

func (s *state[T]) flush(complete func([]T)) {
	s.timer.Stop()
	complete(s.items)
	s.items = nil
}

func main() {
	start := time.Now()

	q := NewCoalescer[time.Duration](4, time.Second, func(batch []time.Duration) {
		fmt.Printf("at %v: %v\n", time.Since(start), batch)
	})
	defer q.Flush()

	for i := 0; i < 1000; i++ {
		q.Put(time.Since(start))
		time.Sleep(time.Duration(rand.Intn(10)) * 80 * time.Millisecond)
	}
}
