package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonbuf := `{
		"company":"itcast",
		"subject":["go","js"],
		"isOk" :true,
		"price":66.66
	}`
	m := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(jsonbuf), &m)
	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	var slice []interface{}
	fmt.Println(m)
	for _, data := range m {
		switch value := data.(type) {
		case string:
			str = value
			fmt.Println("str=", str)
		case []interface{}:
			slice = value
			fmt.Println("slice=", slice)
		}
	}

	fmt.Println(len(slice))
}
