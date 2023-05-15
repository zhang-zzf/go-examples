package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Worker", id, "started", job)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished", job)
		results <- job * 2
	}
}

func main() {
	const jobNum = 5
	jobs := make(chan int, jobNum)
	results := make(chan int, jobNum)
	// start 3 workers
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	fmt.Println("start all jobs at", time.Now())
	for i := 0; i < jobNum; i++ {
		jobs <- i
	}
	result := 0
	for i := 0; i < jobNum; i++ {
		result += <-results
	}
	fmt.Println("result:", result)
	fmt.Println("finish all jobs at", time.Now())
}
