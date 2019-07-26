package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func Printer(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf(string(s[i]))
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
func person1(ch chan int) {
	Printer("hello")
	ch <- 11
}
func person2(ch chan int) {
	<-ch
	Printer("word")
}
func main() {

	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Printf("ch[%d],len(ch)=%d\n", i, len(ch))
		}
	}()
	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}
}
