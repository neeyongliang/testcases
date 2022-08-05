package main

import (
	"fmt"
)

// Basic sends and receives on channels are blocking. However, we
// can use select with a default clause to implement non-blocking sends,
// receives, and even non-blocking multi-way selects.

func main() {
	messages := make(chan string)
	signals := make(chan string)

	// go func() {
	// 	messages <- "one"
	// }()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	// Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
