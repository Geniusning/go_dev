package main

import "fmt"

type Student struct {
	Name  string
	left  *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}
func main() {
	root := new(Student)
	root.Name = "liu"

	left1 := new(Student)
	left1.Name = "liu1"
	root.left = left1

	right1 := new(Student)
	right1.Name = "liu1"
	root.right = right1

	left2 := new(Student)
	left2.Name = "liu2"
	left1.left = left2

	trans(root)
}
