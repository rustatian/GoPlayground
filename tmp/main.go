package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", &Foo{})
	_ = http.ListenAndServe(":8080", m)
}

type Foo struct{}

func (f *Foo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()

	foo := vals["plugin"]
	fmt.Println(foo)
}
