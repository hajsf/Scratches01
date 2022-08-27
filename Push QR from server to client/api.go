package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mux sync.Mutex
)

type DataPasser struct {
	logs chan string
	sema chan struct{} // To control maximum allowed clients connections
}

func (p *DataPasser) handleHello(w http.ResponseWriter, r *http.Request) {
	//setupCORS(&w, r)
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	p.sema <- struct{}{}
	// defer func() { <-p.sema }()
	fmt.Println("from here connection number")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Internal error", 500)
		return
	}

	go Connect()

	for {
		select {
		case c := <-p.logs:
			switch c {
			case "Can not connect with WhatApp server, try again later":
				w.Header().Set("Content-Type", "text/html; charset=ascii")
				fmt.Fprintf(w, "Hello man..")
			default:
				w.Header().Set("Content-Type", "text/event-stream")

				fmt.Println("recieved")
				mux.Lock()
				//counter++
				//c := counter
				mux.Unlock()
				fmt.Fprintf(w, "data: %v\n\n", c)
				flusher.Flush()
			}

		case <-r.Context().Done():
			//<-p.sema
			fmt.Println("Connection closed")
			return
		}
	}
}
