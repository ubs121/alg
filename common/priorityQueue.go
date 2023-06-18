package common

import (
	"container/heap"
	"fmt"
	"testing"
)

// PQItem represents an item in the priority queue.
type PQItem[T any] struct {
	Value    T   // Value of the item
	Priority int // Priority of the item
	Index    int // Index of the item in the heap
}

// PriorityQueue represents a priority queue.
type PriorityQueue[T any] []*PQItem[T]

// Len returns the length of the priority queue.
func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

// Less checks if the item at index i has higher priority than the item at index j.
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the items at indexes i and j.
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem[T])
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority from the priority queue.
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // Mark item as removed
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority of an item in the priority queue.
func (pq *PriorityQueue[T]) Update(item *PQItem[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func TestIntHeap(t *testing.T) {
	// initial values
	h := &IntHeap{2, 1, 5}
	heap.Init(h)

	// push 3
	heap.Push(h, 3)

	fmt.Printf("minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // h.Pop() gives different result !!!
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
