package main

import "fmt"

func main() {

	var a1 = [2][2]int{{1, 2}, {2, 4}}

	for row, v := range a1 {
		for col, v2 := range v {
			fmt.Printf("(%d%d)=%d\n", row, col, v2)
		}
	}
}
