package main

import (
	"fmt"
	"sort"
)

func modify(a map[string]map[string]string) {
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}
	a["zhangsan"]["password"] = "123"
	a["zhangsan"]["nickname"] = "pangopang"
	return
}
func testMap() {
	a := make(map[string]map[string]string, 100)
	a["zhangsan"] = make(map[string]string)
	a["zhangsan"]["age"] = "12"
	modify(a)
	fmt.Println(a)
}
func testMap1() {
	a := make(map[string]map[string]string, 20)
	a["key1"] = make(map[string]string)
	a["key1"]["name1"] = "one"
	a["key1"]["name2"] = "two"
	a["key1"]["name3"] = "three"

	a["key2"] = make(map[string]string)
	a["key2"]["name1"] = "one"
	a["key2"]["name2"] = "two"
	a["key2"]["name3"] = "three"
	for index, val := range a {
		fmt.Println(index)
		for index, val2 := range val {
			fmt.Println("\t", index, val2)
		}
	}
}
func testMap2() {
	var a []map[string]string
	a = make([]map[string]string, 5)
	if a[0] == nil {
		a[0] = make(map[string]string)
	}
	a[0]["name"] = "liuning"
	fmt.Println(a)
}
func mapSort() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[1] = 102
	a[5] = 104
	a[6] = 105
	a[2] = 106
	a[8] = 101

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}
func main() {
	mapSort()
}
