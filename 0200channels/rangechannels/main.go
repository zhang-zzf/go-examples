package main

func main() {
	queue := make(chan string, 2)
	done := make(chan bool)
	queue <- "hello"
	queue <- "world"

	go func() {
		for v := range queue {
			println("receive:", v)
		}
		done <- true

	}()

	queue <- "2"
	queue <- "1"
	close(queue)
	<-done
}
