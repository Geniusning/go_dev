package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type typeCount struct {
	chCount    int
	numCount   int
	spaceCount int
	otherCount int
}

func main() {
	// file, err := os.OpenFile("D:/test.log", os.O_CREATE|os.O_WRONLY, 0664)
	// if err != nil {
	// 	fmt.Println("open file failed :", err)
	// 	return
	// }
	// fmt.Fprintf(file, "do more time\n")
	// file.Close()
	file, err := os.Open("D:/test.log")
	if err != nil {
		fmt.Println("read fail,err", err)
		return
	}
	defer file.Close()

	var count typeCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("reader fail err:", err)
			break
		}

		runeArray := []rune(str)
		for _, v := range runeArray {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.chCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			case v >= '0' && v <= '9':
				count.numCount++
			default:
				count.otherCount++
			}
		}
		fmt.Println(str)
	}
	fmt.Println("str has char count:", count.chCount)
}
