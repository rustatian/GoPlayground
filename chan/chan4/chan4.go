package main

import (
	"context"
	"time"
)

func main() {
	s := NewStr()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	//s.put()
	//for i := 0; i < 1000; i++ {
	s.foo(ctx)
	//}
	cancel()

	s.put()
	s.put()
	s.put()

	s.foo(context.Background())
}

type str struct {
	v chan struct{}
}

func NewStr() *str {
	return &str{v: make(chan struct{}, 5)}
}

func (s *str) put() {
	s.v <- struct{}{}
}

func (s *str) foo(ctx context.Context) {
	select {
	case <-ctx.Done():
		return
	case <-s.v:
		return
	}
}
