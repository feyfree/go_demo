package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello h2c")
	})
	s := &http.Server{
		Addr:    ":8972",
		Handler: mux,
	}
	http2.ConfigureServer(s, &http2.Server{})
	log.Fatal(s.ListenAndServe())
}
