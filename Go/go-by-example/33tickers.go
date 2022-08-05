package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the
// future - tickers are for when you want to do something
// repeatedly at regular intervals. Here’s an example of a
// ticker that ticks periodically until we stop it.

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
