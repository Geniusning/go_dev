package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := Student{Person{"liuning", 12, "1"}, 111, "shenzhen"}
	fmt.Printf("s:%v", s)
	fmt.Println(s.name)
}
