package main

import "time"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			jobId, more := <-jobs
			if more {
				println("exec job: ", jobId)
				time.Sleep(time.Second * 3)
			} else {
				println("exec all jobs")
				done <- true
				break
			}
		}
	}()
	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs)
	// wait for all jobs executed
	<-done
}
