package main

import (
	"fmt"
	"time"
)

// Here’s the worker, of which we’ll run several
// concurrent instances. These workers will
// receive work on the jobs channel and send the
// corresponding results on results. We’ll
// sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// In order to use our pool of workers we need to send them
	// work and collect their results. We make 2 channels for this.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// This starts up 3 workers, initially blocked because there
	// are no jobs yet.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	// Finally we collect all the results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
