package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	_, err := fmt.Fprintf(ctx, "Hello World!")
	if err != nil {
		panic(err)
	}
}

func main() {
	// pass plain function to fasthttp
	_ = fasthttp.ListenAndServe(":8000", fastHTTPHandler)
}
