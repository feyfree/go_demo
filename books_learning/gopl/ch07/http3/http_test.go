// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 194.

// Http3 is an e-commerce server that registers the /list and /price
// endpoints by calling (*http.ServeMux).Handle.
package http3

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

//!+main

func TestServerMux(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	// http.HandlerFunc(db.list) is a conversion, not a function call
	// 是转换， 而不是函数调用
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

//!-main

/*
//!+handlerfunc
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
//!-handlerfunc
*/
