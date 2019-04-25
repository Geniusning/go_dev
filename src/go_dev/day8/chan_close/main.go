package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("reveive job", job)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("send job", i)
	}
	close(jobs)
	fmt.Println("send all jobs")

	<-done
}
