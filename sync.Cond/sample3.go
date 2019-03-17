package main

import (
	"fmt"
	"sync"
	"time"
)

type PartitionLocker struct {
	c *sync.Cond
	l sync.Locker
	s map[string]struct{}
}

func NewPartitionLocker(l sync.Locker) *PartitionLocker {
	return &PartitionLocker{c: sync.NewCond(l), l: l, s: make(map[string]struct{})}
}

func (p *PartitionLocker) locked(id string) (ok bool) {
	_, ok = p.s[id]
	return
}

func (p *PartitionLocker) Lock(id string) {
	p.l.Lock()
	defer p.l.Unlock()
	for p.locked(id) {
		p.c.Wait()
	}
	p.s[id] = struct{}{}
	return
}

func (s *PartitionLocker) Unlock(id string) {
	s.l.Lock()
	defer s.l.Unlock()
	delete(s.s, id)
	s.c.Broadcast()
}

func main() {
	var (
		mu  sync.Mutex
		wg  sync.WaitGroup
		ids = []string{
			"red",
			"blue",
			"yellow",
		}
	)

	pl := NewPartitionLocker(&mu)
	for i := 0; i < 90; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			pl.Lock(ids[i%3])
			time.Sleep(time.Second)
			fmt.Println(i, ids[i%3])
			pl.Unlock(ids[i%3])
		}()
	}
	wg.Wait()
}
