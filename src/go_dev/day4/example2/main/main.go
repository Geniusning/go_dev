package main

import (
	"fmt"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println(a)
	return
}
func calc(n int) int {
	if n == 1 {
		return 1
	}

	return calc(n-1) * n

}
func main() {
	c := calc(5)
	fmt.Println(c)
}
