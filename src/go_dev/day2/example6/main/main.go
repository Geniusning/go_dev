package main

import "fmt"

//计算水仙花数
func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		sum += s
	}
	return sum
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	s := sum(n)
	fmt.Println(s)
}
