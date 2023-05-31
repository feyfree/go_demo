package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

// 1. 创建CA证书
// openssl genrsa -out rootCA.key 2048
// openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem

// 2. 创建证书
// openssl genrsa -out server.key 2048
// openssl req -new -key server.key -out server.csr
// openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	dir, _ := os.Getwd()
	path := "http_demo/http1.1/tls"
	crt := dir + "/" + path + "/" + "server.crt"
	key := dir + "/" + path + "/" + "server.key"
	log.Fatal(http.ListenAndServeTLS(":443", crt, key, nil))
}
