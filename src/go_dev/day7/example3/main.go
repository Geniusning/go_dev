package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("reader is failer,err:", err)
	}

	fmt.Printf("read str succ, %s\n", str)
}
