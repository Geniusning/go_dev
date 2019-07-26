package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL="+r.URL.String())
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	//read local file path
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe("127.0.0.1:8089", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is sayhello")
}
