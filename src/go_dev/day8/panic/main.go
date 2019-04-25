package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err:", err)
		}
	}()
	var m map[string]string
	m["name"] = "liuning"
}

func calc() {
	for {
		fmt.Println("i am calcing")
		time.Sleep(time.Second)
	}
}
func main() {

	go test()
	for i := 0; i < 2; i++ {
		go calc()
	}
	time.Sleep(time.Second * 10)
}
