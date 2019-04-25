package main

import "fmt"

func send(ch chan int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	var a struct{}
	exitChan <- a
}

func recv(ch chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("is all")
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}

func main() {
	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go func() {
		go send(ch, exitChan)
	}()

	go func() {
		go recv(ch, exitChan)
	}()

	<-exitChan
	<-exitChan
}
