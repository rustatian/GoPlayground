package main

import "net/http"

func main() {
	http.ListenAndServe("localhost:2112", nil)
}
