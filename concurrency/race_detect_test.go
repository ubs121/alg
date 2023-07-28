package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

var counter int

func increment() {
	counter++
}

// Use "go test -race"
func TestRaceDetection(t *testing.T) {
	var wg sync.WaitGroup
	const goroutines = 100

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
