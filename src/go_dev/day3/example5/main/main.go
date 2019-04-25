package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)

	resultStr, error := strconv.Atoi(str)
	if error != nil {
		fmt.Println("can not convert to int")
	}
	fmt.Println("result_str", resultStr)
}
