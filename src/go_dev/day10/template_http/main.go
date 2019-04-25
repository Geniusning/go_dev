package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var myTemplate *template.Template

type Student struct {
	Name  string
	Age   int
	Title string
}
type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("call by template")
	p.output += string(b)
	return len(b), nil
}
func userInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle userInfo")
	p := Student{Name: "liuning", Age: 13, Title: "GOOD"}
	resultWriter := &Result{}
	myTemplate.Execute(resultWriter, p)
	fmt.Println("template render data :", resultWriter.output)

}
func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("template is failed,err:", err)
		return
	}
	return
}
func main() {
	initTemplate("C:/Users/Administrator/go/src/go_dev/day10/template_http/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("127.0.0.1:9987", nil)
	if err != nil {
		fmt.Println("listen http failed,err:", err)
	}
}
