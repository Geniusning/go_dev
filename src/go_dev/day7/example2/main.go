package main

import "fmt"

type student struct {
	name  string
	age   int
	score float32
}

func main() {
	var str = "stu01 18 89.2"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.name, &stu.age, &stu.score)
	fmt.Println(stu)

}
