package main

import (
	"fmt"
	"html/template"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Title string
}

func main() {
	t, err := template.ParseFiles("C:/Users/Administrator/go/src/go_dev/day10/template/index.html")
	if err != nil {
		fmt.Println("template is failed,err:", err)
		return
	}

	p := Student{Name: "liuning", Age: 13, Title: "GOOD"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("there are an error:", err.Error())
	}
}
