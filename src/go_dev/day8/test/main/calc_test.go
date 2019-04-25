package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(1, 1)
	if r != 3 {
		fmt.Printf("err:1+2=%d,actual=%d\n", 3, r)
		return
	}

	fmt.Println("test succ")
}
