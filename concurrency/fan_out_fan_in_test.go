package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func workerFOFI(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		// Simulating some work
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Perform the work and send the result
		result := job * 2
		results <- result
		fmt.Printf("Worker %d processed job %d, result: %d\n", id, job, result)
	}
}

func TestFOFI(t *testing.T) {
	numWorkers := 5
	numJobs := 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Generate jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Create wait group to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Start workers, distribute works among multiple workers (fan-out)
	for i := 1; i <= numWorkers; i++ {
		go func(workerID int) {
			workerFOFI(workerID, jobs, results)
			wg.Done()
		}(i)
	}

	// Close the results channel when all workers have finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results from the results channel (fan-in)
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}

	fmt.Println("All jobs have been processed.")
}
