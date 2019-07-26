package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string   `json:"company"`
	Subject []string `json:"subject"`
	IsOk    bool     `json:"isOk"`
	Price   float64  `json:"price"`
}

func main() {
	jsonbuf := `{
		"company":"itcast",
		"subject":["go","js"],
		"isOk" :true,
		"price":66.66
	}`
	var t IT
	err := json.Unmarshal([]byte(jsonbuf), &t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", t)
}
