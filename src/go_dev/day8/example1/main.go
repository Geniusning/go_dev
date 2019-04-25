package main

import "fmt"

type student struct {
	name string
}

func main() {
	var intChan chan interface{}
	intChan = make(chan interface{}, 10)
	stu := student{name: "liuning"}
	intChan <- &stu

	var stu01 interface{}
	stu01 = <-intChan

	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("convert faild")
		return
	}
	fmt.Println(stu02)
}
