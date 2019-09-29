package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)



// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "Hello")
}

func main()  {
	// pass plain function to fasthttp
	_ = fasthttp.ListenAndServe(":7878", fastHTTPHandler)
}

