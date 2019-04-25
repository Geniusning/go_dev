package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	socre float32
	next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(p)
		p = p.next
	}
}
func tailInsert(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {
		var stu = Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			socre: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
}
func addNode(p *Student, newNode *Student) {
	for p != nil {
		if p.Name == "stu2" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}
func main() {
	var head Student
	head.Name = "liu"
	head.Age = 12
	head.socre = 80

	tailInsert(&head)
	var newNode *Student = new(Student)
	newNode.Name = "liu124"
	newNode.Age = 1232
	newNode.socre = 801
	addNode(&head, newNode)
	trans(&head)
}
