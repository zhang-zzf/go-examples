package main

func main() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		println("received:", msg)
	default:
		println("no message received")
	}

	select {
	case messages <- "Hello, World":
		println("massage sent")
	default:
		println("no message send")
	}

	select {
	case <-messages:
	case <-signals:
	default:
		println("no activity")
	}
}
