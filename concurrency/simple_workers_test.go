package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func simpleWorker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Perform some work
	fmt.Printf("Worker %d started\n", id)
	// ...

	// Send result to the channel
	ch <- id
	fmt.Printf("Worker %d finished\n", id)
}

func TestSimpleWorker(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)

	// Number of workers
	numWorkers := 3

	// Add the number of workers to the WaitGroup
	wg.Add(numWorkers)

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		go simpleWorker(i, ch, &wg)
	}

	// Close the channel when all workers have finished
	go func() {
		wg.Wait() // Wait for all workers to finish
		close(ch) // Close the channel
	}()

	// Receive results from the channel
	for result := range ch {
		fmt.Printf("Received result from Worker %d\n", result)
	}

	fmt.Println("All workers have finished.")
}
