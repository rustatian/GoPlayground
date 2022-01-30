package main

import (
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAll(resp.Body)
	_ = resp.Body.Close()
}
