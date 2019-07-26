package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createRand(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num > 1000 {
			break
		}
	}
	*p = num
}
func getNum(randSlice []int, p *int) {
	randSlice[0] = *p / 1000
	randSlice[1] = (*p % 1000) / 100
	randSlice[2] = (*p % 100) / 10
	randSlice[3] = *p % 10
}

func onGame(randSlice []int) {
	var userNum int
	for {
		for {
			fmt.Printf("请输入四位数：")
			fmt.Scan(&userNum)
			if userNum > 999 && userNum < 10000 {
				break
			}
			fmt.Println("请输入有效的四位数！")
		}
		keySlice := make([]int, 4)
		n := 0
		getNum(keySlice, &userNum)
		for i := 0; i < len(keySlice); i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d个数大一点\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d个数小一点\n", i+1)
			} else {
				fmt.Printf("第%d个数正确\n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("数字猜对啦")
			break
		}
	}
}
func main() {
	var rankNum int
	createRand(&rankNum)
	randSlice := make([]int, 4)

	getNum(randSlice, &rankNum)

	onGame(randSlice)
	fmt.Println("randSlice=", randSlice)
}
