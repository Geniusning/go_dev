package main

import "fmt"

type addFun func(int, int) int

func add(a, b int) int {
	return a + b
}

func operate(c addFun, i int, j int) int {
	return c(i, j)
}

func main() {
	c := add
	sum := operate(c, 1, 1)
	fmt.Println(sum)
}
