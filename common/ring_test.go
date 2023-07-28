package common

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	// Create a new ring with 5 elements
	r := ring.New(5)

	// Populate the ring with values
	for i := 1; i <= 5; i++ {
		r.Value = i
		r = r.Next()
	}

	// Traverse the ring and verify the values
	r.Do(func(value interface{}) {
		fmt.Print(value, ",")
	})
	fmt.Println()

	// Move the ring pointer by 2 positions
	r = r.Move(2)

	// Traverse the ring and verify the values
	r.Do(func(value interface{}) {
		fmt.Print(value, ",")
	})
	fmt.Println()

	// Get the value at the current position
	value := r.Value.(int)
	expectedValue := 3

	// Check if the value matches the expected value
	if value != expectedValue {
		t.Errorf("Expected value %d, but got %d", expectedValue, value)
	}
}
