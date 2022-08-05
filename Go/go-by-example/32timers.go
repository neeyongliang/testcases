package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// The <-timer1.C blocks on the timer’s channel C until it sends a value
	// indicating that the timer expired.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer("Timer 2 expired")
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
