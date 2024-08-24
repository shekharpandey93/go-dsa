package main

import (
	"context"
	"fmt"
	"time"
)

func printVal(ct context.Context, id int) {
	for {
		select {
		case <-ct.Done():
			fmt.Printf("Worker %d: Stopping work due to cancellation\n", id)
			return
		default:
			// Simulate work
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(500 * time.Millisecond) // Simulate work by sleeping
		}

	}
}

func main() {
	context, cancel := context.WithCancel(context.Background())
	// Start multiple workers
	for i := 1; i <= 3; i++ {
		go printVal(context, i)
	}
	// Let the workers run for a while
	time.Sleep(2 * time.Second)

	// Cancel the context to signal all workers to stop
	fmt.Println("Main: Cancelling context...")
	cancel()

	//// Give the workers time to clean up
	//time.Sleep(1 * time.Second)
	//
	fmt.Println("Main: All workers stopped.")

}
