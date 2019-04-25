package main

import (
	"fmt"
	"go_dev/day7/example1/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var ints []*balance.Instances
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstace(host, 8080)
		ints = append(ints, one)
	}
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		ints, err := balance.DoBalance(balanceName, ints)
		if err != nil {
			fmt.Println("do balance err；", err)
			continue
		}

		fmt.Println(ints)
		time.Sleep(time.Second)
	}

}
