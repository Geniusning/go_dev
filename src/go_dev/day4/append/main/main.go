package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice := arr[1:]
	// fmt.Printf("arr=%p,slice=%p\n", &arr[1], slice)
	// slice = append(slice, 10)
	// slice = append(slice, 10)
	// slice = append(slice, 10)
	// fmt.Println(slice)
	// fmt.Printf("arr=%p,slice=%p\n", &arr[1], slice)

	//--------------------------------------------
	str := "你sringiag"
	s1 := []rune(str)

	s1[0] = '0'
	as := string(s1)
	fmt.Println(as)
}
