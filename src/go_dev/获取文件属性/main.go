package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("write args wrong,please input 2 param")
		return
	}

	fileInfo, err := os.Stat(list[1])
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("fileName=", fileInfo.Name())
	fmt.Println("size=", fileInfo.Size())
}
