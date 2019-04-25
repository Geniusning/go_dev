package main

import "fmt"

func modify(p *int) {

	*p = 100

}

func main() {
	var a int = 10
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

}
