package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type IT struct {
	Company string   `json:"company"`
	Subject []string `json:"subject"`
	IsOk    bool     `json:",string"`
	Price   float64  `json:",string"`
}

func main() {
	buf := "a1c akc apc amc dds"

	reg := regexp.MustCompile(`a.c`)
	if reg == nil {
		fmt.Println("mustCompile failed")
		return
	}
	it := IT{"itcast", []string{"go", "js"}, true, 66.666}
	value, err := json.MarshalIndent(it, "", " ")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(string(value))
	slice := reg.FindAllString(buf, 1)
	fmt.Println(slice)

}
