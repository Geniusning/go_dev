package main

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("counting")
}

func main() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
		time.Sleep(time.Second)
	}
}
