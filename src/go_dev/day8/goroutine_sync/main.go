package main

import (
	"fmt"
)

func calc(taskChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			resultChan <- v
		}
	}
	exitChan <- true
}

func main() {
	taskChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 2; i < 10000; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(taskChan, resultChan, exitChan)
	}
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
	}()
	for v := range resultChan {
		fmt.Println(v)
	}
}
