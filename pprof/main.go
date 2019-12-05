package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	_ "net/http/pprof"
)



// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "Hello World!")
}

func main()  {
	go func() {
		log.Println(http.ListenAndServe("localhost:6061", nil))
	}()

	// pass plain function to fasthttp
	_ = fasthttp.ListenAndServe(":8000", fastHTTPHandler)
}