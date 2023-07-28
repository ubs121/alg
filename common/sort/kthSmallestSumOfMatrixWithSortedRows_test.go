package sort

import (
	"container/heap"
	"fmt"
	"testing"
)

// https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/
// You are given an m x n matrix mat that has its rows sorted in non-decreasing order and an integer k.
// You are allowed to choose exactly one element from each row to form an array.
// Return kth smallest sum among all possible arrays.
func kthSmallest(matrix [][]int, k int) int {
	sums := matrix[0] // first row

	// push all the remaining rows
	for i := 1; i < len(matrix); i++ {
		sums = updateSmallestSums(sums, matrix[i], k)
	}

	// 'sums' now should contain the k smallest sums, and the k-th smallest sum is returned
	return sums[k-1]
}

func updateSmallestSums(preSums []int, currentRow []int, k int) []int {

	// create a min heap
	h := new(MinHeap)
	heap.Init(h)

	// push the first pairs of sums into the heap
	// at maximum, take 'k' elements
	for i := 0; i < len(preSums) && i < k; i++ {
		heap.Push(h, sumPair{0, preSums[i], currentRow[0]})
	}

	// sort & collect minimum elements
	minSums := make([]int, 0, k)
	for k > 0 && h.Len() > 0 {
		pair := heap.Pop(h).(sumPair)                // get the minimum element
		minSums = append(minSums, pair.sum+pair.val) // add to the result

		// If the popped element is not the last element of the row, push the next element of the same row to the heap
		if pair.col+1 < len(currentRow) {
			heap.Push(h, sumPair{pair.col + 1, pair.sum, currentRow[pair.col+1]})
		}

		k--
	}
	return minSums
}

func TestKthSmallest(t *testing.T) {
	testCases := []struct {
		data [][]int
		k    int
		exp  int
	}{
		{
			[][]int{
				{1, 3, 11},
				{2, 4, 6},
			},
			5, 7,
		},
		{
			[][]int{
				{1, 3, 11},
				{2, 4, 6},
			},
			1, 3,
		},
		{
			[][]int{
				{1, 3, 11},
				{2, 4, 6},
			},
			9, 17,
		},
		{
			[][]int{
				{1, 10, 10},
				{1, 4, 5},
				{2, 3, 6},
			},
			7, 9,
		},
		{
			[][]int{
				{1, 10, 10},
				{1, 4, 5},
				{2, 3, 6},
			},
			1, 4,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := kthSmallest(tc.data, tc.k)
			if got != tc.exp {
				t.Errorf("test case %d: exp %d got %d", i, tc.exp, got)
			}
		})
	}
}

type sumPair struct {
	col int // column index
	sum int // previuos sum
	val int // new value
}

type MinHeap []sumPair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum+h[i].val < h[j].sum+h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(sumPair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
