package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hellp")
	fmt.Fprintf(w, "hello")
}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
	fmt.Fprintf(w, "欢迎来到登录页面")
}
func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", Login)
	err := http.ListenAndServe("127.0.0.1:9999", nil)
	if err != nil {
		fmt.Println("http listen failed", err)
	}
}
