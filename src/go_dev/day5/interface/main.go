package main

import "fmt"

type Test interface {
	print()
}
type student struct {
	name  string
	age   int
	score int
}

func (p *student) print() {
	fmt.Println("name is ", p.name)
	fmt.Println("name is ", p.age)
	fmt.Println("name is ", p.score)
}

func main() {
	var t Test
	var s student = student{
		name:  "liuning",
		age:   10,
		score: 12,
	}
	t = &s
	t.print()
}
