package main

import (
	"fmt"
	"go_dev/day2/example2/add"
)

func init() {
	add.Test()
}
func main() {
	fmt.Println("reslut", add.Name)
	fmt.Println("reslut", add.Age)
}
