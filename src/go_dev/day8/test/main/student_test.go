package main

import (
	"fmt"
	"testing"
)

func TestStudent(t *testing.T) {
	stu := &Student{
		name: "liuning",
		age:  12,
	}
	err := stu.Save()
	if err != nil {
		t.Fatalf("save failed,err:%v\n", err)
	}

	stu2 := &Student{
		name: "liuning",
		age:  12,
	}
	err1 := stu2.Load()
	if err != nil {
		t.Fatalf("save failed,err:%v\n", err1)
	}

	if stu.name != stu2.name {
		fmt.Println("stu1.name is equal")
	}
	fmt.Println("test student succ")
}
