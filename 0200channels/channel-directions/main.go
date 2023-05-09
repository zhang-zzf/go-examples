package main

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// invalid operation: pings <- "aaa" (send to the receive-only type <-chan string)
	// pings <- "aaa"
	pongs <- <-pings
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hello,World.")
	pong(pings, pongs)
	println(<-pongs)

}
