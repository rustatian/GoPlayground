package main

import (
	"log"
	"net/http"
	"sync"
)

type user struct {
	name string
	age  uint8
	port string
}

type T1 string
type T2 string

func (T1) init() *user {
	u1 := user{
		port: "8080",
		name: "Andrei",
		age:  48,
	}
	return &u1
}

func (T2) init() *user {
	u2 := user{
		port: "8081",
		name: "Nicolay",
		age:  24,
	}
	return &u2
}

func main() {

	var users = make([]*user, 2)

	var t1 T1 // Doesn't work
	users[1] = t1.init()

	//var t2 T2  // Works
	//users[0] = t2.init()

	serve(users...)

}

func serve(us ...*user) {

	wg := &sync.WaitGroup{}

	for _, u := range us {

		if u == nil {
			continue
		}

		wg.Add(1)
		go func() {
			log.Printf("Serving user with name %v of age %v", u.name, u.age)
			log.Fatal(http.ListenAndServe(":"+u.port, u))
			wg.Done()
		}()

	}

	wg.Wait()

}

func (u user) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hey " + u.name))
}
