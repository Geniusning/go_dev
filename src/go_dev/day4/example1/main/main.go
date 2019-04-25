package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func multis() { //99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
func qiuYinZi() { //求完数
	for i := 1; i < 1000; i++ {
		var sum int
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}
		if i == sum {
			fmt.Println("完数", sum)
		}
	}
}
func huiWen() { //求回文
	var s string
	fmt.Scanf("%s", &s)
	arr := strings.Split(s, "")
	var judgeB bool
	for i := 0; i < len(arr); i++ {
		if i == len(arr)/2 {
			break
		}
		if arr[i] == arr[len(arr)-i-1] {
			judgeB = true
		} else {
			judgeB = false
		}
	}
	if judgeB {
		fmt.Println("is hui wen")
	} else {
		fmt.Println("is not hui wen")
	}
}
func count(str string) (wordCount, spaceCount, numberCount, otherCount int) { //计数
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}
func multiString(str1, str2 string) (result string) { //大数相加
	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return result
	}
	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c2 := str2[index2] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}
	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}

	splicResult := strings.Split(string(result), "+")
	if len(splicResult) != 2 {
		fmt.Println("please input Tab")
		return
	}

	strings1 := strings.TrimSpace(splicResult[0])
	strings2 := strings.TrimSpace(splicResult[1])

	fmt.Println(multiString(strings1, strings2))

}
