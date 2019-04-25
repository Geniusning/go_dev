package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//互斥锁
var lock sync.Mutex
var rwLock sync.RWMutex

func main() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32
	a[1] = 102
	a[5] = 104
	a[6] = 105
	a[2] = 106
	a[8] = 101
	for i := 0; i < 3; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			a[1] = rand.Intn(10)
			atomic.AddInt32(&count, 1)
			time.Sleep(10 * time.Millisecond)
			rwLock.Unlock()
		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			a[1] = rand.Intn(10)
			atomic.AddInt32(&count, 1)
			rwLock.Unlock()
		}(a)
	}

	time.Sleep(time.Second * 3)
	rwLock.Lock()
	fmt.Println(atomic.LoadInt32(&count))
	rwLock.Unlock()

}
