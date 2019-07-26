package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var muxHandlerMap map[string]func(w http.ResponseWriter, r *http.Request)

type myMuxHandle struct{}

func (*myMuxHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := muxHandlerMap[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "Url="+r.URL.String())
}
func main() {
	server := http.Server{
		Addr:        ":8080",
		ReadTimeout: 5 * time.Second,
		Handler:     &myMuxHandle{},
	}

	muxHandlerMap = make(map[string]func(w http.ResponseWriter, r *http.Request))
	muxHandlerMap["/hello"] = onSayHello
	muxHandlerMap["/hi"] = onSayHi

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func onSayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is sayHello")
}

func onSayHi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is sayHi")
}
