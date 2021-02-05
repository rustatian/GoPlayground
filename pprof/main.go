package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/valyala/fasthttp"
)

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprintf(ctx, "Hello World!")
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6063", nil))
	}()

	for i := 0; i <= 1000; i++ {
		time.Sleep(time.Second * 1)
	}

	// pass plain function to fasthttp
	_ = fasthttp.ListenAndServe(":8000", fastHTTPHandler)

}
