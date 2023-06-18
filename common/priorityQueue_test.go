package common

import (
	"container/heap"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// Create a priority queue of integers
	pq := make(PriorityQueue[int], 0)

	// Push items into the priority queue
	item1 := &PQItem[int]{Value: 42, Priority: 2}
	item2 := &PQItem[int]{Value: 13, Priority: 1}
	item3 := &PQItem[int]{Value: 7, Priority: 3}
	heap.Push(&pq, item1)
	heap.Push(&pq, item2)
	heap.Push(&pq, item3)

	// Expected output in order of priority
	expected := []struct {
		Value    int
		Priority int
	}{
		{Value: 13, Priority: 1},
		{Value: 42, Priority: 2},
		{Value: 7, Priority: 3},
	}

	// Pop items from the priority queue and compare with expected output
	for _, ex := range expected {
		item := heap.Pop(&pq).(*PQItem[int])
		if item.Value != ex.Value || item.Priority != ex.Priority {
			t.Errorf("Expected Value: %d, Priority: %d. Got Value: %d, Priority: %d",
				ex.Value, ex.Priority, item.Value, item.Priority)
		}
	}
}
