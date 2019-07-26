package main

import (
	"fmt"
	"runtime"
)

func test(ch chan string) {
	defer fmt.Println("cccccccccc")
	ch <- "abd"
	fmt.Println("bbbbbbbbbbbbbb")
}
func main() {
	n := runtime.GOMAXPROCS(10)
	fmt.Println(n)
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
