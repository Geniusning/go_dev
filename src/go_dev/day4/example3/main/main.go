package main

import (
	"fmt"
	"strings"
)

func adder() func(a int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	f := makeSuffix(".txt")
	fmt.Println(f("123"))
	fmt.Println(f("test"))
	fmt.Println(f("test.txt"))
}
