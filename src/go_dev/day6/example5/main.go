package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	name string
	age  int
}

type Book struct {
	name  string
	count int
}

type studentArray []Student

func (p studentArray) Len() int {
	return len(p)
}
func (p studentArray) Less(i, j int) bool {
	return p[i].name > p[j].name
}
func (p studentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	var stus studentArray
	for i := 0; i < 10; i++ {
		stu := Student{
			name: fmt.Sprintf("stu%d", rand.Intn(100)),
			age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Println("\n\n")

	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}
}
