package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "liu",
		Age:   3124,
		Score: 324,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed")
	}
	fmt.Println(string(data))
}
