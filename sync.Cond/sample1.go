package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Record struct {
	sync.Mutex

	buf  string
	cond *sync.Cond

	writers []io.Writer
}

func NewRecord(writers ...io.Writer) *Record {
	r := &Record{writers: writers}
	r.cond = sync.NewCond(r)
	return r
}

func (r *Record) Prompt() {
	for {

		fmt.Printf(":> ")
		var s string
		fmt.Scanf("%s", &s)
		if s == "exit" {
			os.Exit(0)
			//r.cond.Broadcast()
		}

		r.Lock()
		r.buf = s
		r.Unlock()

		r.cond.Signal()
	}
}

func (r *Record) Start() error {
	f := func(w io.Writer) {
		for {
			r.Lock()
			r.cond.Wait()
			fmt.Fprintf(w, "%s\n", r.buf)
			r.Unlock()

		}
	}
	for i := range r.writers {
		go f(r.writers[i])
	}
	return nil
}

func main() {
	f, err := os.Create("./cond.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := NewRecord(f)
	err = r.Start()
	if err != nil {

	}
	r.Prompt()
}
