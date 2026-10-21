package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(10)

	for i := range 100 {
		g.Go(func() error {
			return process(ctx, i)
		})
	}

	if err := g.Wait(); err != nil {
		// handle error
		if err != nil {
			panic(err)
		}
	}
}

type ConcreteT struct{}

func (c ConcreteT) Foo() ConcreteT {
	return c
}

func process(ctx context.Context, i int) error {
	// do something
	println(i)
	return nil
}

type Generic interface {
	int | string
}

func processt[T any](ctx context.Context, i T) error {
	// do something
	println(i)
	return nil
}

type Adder[T Adder[T]] interface {
	Add(T) T
}
