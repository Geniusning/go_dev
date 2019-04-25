package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
	}
}
func main() {
	ch := make(chan int, 10)
	go write(ch)
	go read(ch)
	time.Sleep(time.Second * 10)
}
