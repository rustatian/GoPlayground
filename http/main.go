package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)



// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello")
}

func main()  {
	// pass plain function to fasthttp
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}

