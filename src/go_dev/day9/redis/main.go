package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("conn redis faied,", err)
		return
	}
	defer c.Close()
}
