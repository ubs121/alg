package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncCond(t *testing.T) {
	var cond sync.Cond
	var wg sync.WaitGroup

	// Initialize the conditional variable with a Locker
	cond.L = new(sync.Mutex)

	// Number of goroutines
	numGoroutines := 5

	// Add the number of goroutines to the WaitGroup
	wg.Add(numGoroutines)

	// Start the goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			// Acquire the lock associated with the conditional variable
			cond.L.Lock()
			defer cond.L.Unlock()

			// Wait until a condition is met
			fmt.Printf("Goroutine %d waiting...\n", id)
			cond.Wait()

			// Execute after the condition is signaled
			fmt.Printf("Goroutine %d woken up!\n", id)

			// Signal completion to the WaitGroup
			wg.Done()
		}(i)
	}

	// Allow some time for the goroutines to start
	time.Sleep(time.Second)

	// Signal the condition to wake up the waiting goroutines
	fmt.Println("Wake up all...")
	cond.Broadcast()

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All goroutines completed.")
}
