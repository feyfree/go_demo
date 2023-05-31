package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"os"
	"time"
)

// 1. 创建CA证书
// openssl genrsa -out rootCA.key 2048
// openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem

// 2. 创建证书
// openssl genrsa -out server.key 2048
// openssl req -new -key server.key -out server.csr
// openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500

const idleTimeout = 5 * time.Minute
const activeTimeout = 10 * time.Minute

func main() {
	var srv http.Server
	//http2.VerboseLogs = true
	srv.Addr = ":8972"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http2"))
	})
	http2.ConfigureServer(&srv, &http2.Server{})
	dir, _ := os.Getwd()
	path := "http_demo/http2/tls"
	crt := dir + "/" + path + "/" + "server.crt"
	key := dir + "/" + path + "/" + "server.key"
	go func() {
		log.Fatal(srv.ListenAndServeTLS(crt, key))
	}()
	select {}

}
