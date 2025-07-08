package async

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this goroutine is done

	println("Worker", id, "is starting")

	// Simulate some work
	for i := range 100 {
		// Do some work here
		fmt.Print("Worker ", id, " is working on iteration ", i, "\n")
	}
}

func Run_WG() {
	var wg sync.WaitGroup

	// Start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)                   // Increment the WaitGroup counter
		go worker(i, &wg)           // Start a new goroutine
		time.Sleep(2 * time.Second) // Sleep to simulate staggered start
	}

	wg.Wait() // Wait for all workers to finish
	println("All workers have completed their tasks")
}
