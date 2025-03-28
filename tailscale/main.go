package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	"tailscale.com/client/tailscale"
)

func main() {
	s := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: tailscale.GetCertificate,
		},
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "<h1>Hello from Tailscale!</h1> It works.")
		}),
	}
	log.Printf("Running TLS server on :443 ...")
	log.Fatal(s.ListenAndServeTLS("", ""))
}
