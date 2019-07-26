package main

import "fmt"

type FuncTest func(int, int) int

func Add(a int, b int) (result int) {
	result = a + b
	return
}

func Calc(a, b int, Ftext FuncTest) (result int) {
	result = Ftext(a, b)
	return
}
func main() {
	result := Calc(1, 2, Add)
	fmt.Println(result)
}
