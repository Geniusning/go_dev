package main

import (
	"fmt"
	"sort"
)

func testSort() {
	arr := [5]int{3, 4, 24, 5, 1}
	sort.Ints(arr[:])
	fmt.Println(arr)
}
func searchInt() {
	arr := [5]int{3, 4, 24, 5, 1}
	sort.Ints(arr[:])
	index := sort.SearchInts(arr[:], 24)
	fmt.Println(index)
}
func main() {
	searchInt()
}
