package main

import (
	"sync"
	"testing"
)

func TestAppend(t *testing.T) {
	x := make([]string, 0, 6)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(len(z))
	}()
	wg.Wait()
}
