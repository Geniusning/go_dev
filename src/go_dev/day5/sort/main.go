package main

import "fmt"

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j+1] < slice[j] {
				slice[j+1], slice[j] = slice[j], slice[j+1]
			}
		}
	}
}
func isort(slice []int) { //插入排序
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			if slice[j] > slice[j-1] {
				break
			}
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
}
func fastSort() {

}
func main() {
	var arr = [...]int{12, 45, 678, 97, 54, 33, 67, 43}
	isort(arr[:]) //冒泡排序
	fmt.Println(arr)
}
