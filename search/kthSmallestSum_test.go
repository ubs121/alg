package search

import "testing"

// https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/
func kthSmallest(mat [][]int, k int) int {
	// TODO: push into a priority queue, stop when k<i^m element inserted, return kth element
	return 0
}

func TestKthSmallest(t *testing.T) {
	// input format
	type input struct {
		data [][]int
		k    int
	}
	// test cases
	testCases := map[*input]int{
		{
			data: [][]int{{1, 3, 11}, {2, 4, 6}},
			k:    5,
		}: 7,
	}

	for tc, exp := range testCases {
		got := kthSmallest(tc.data, tc.k)
		if got != exp {
			t.Errorf("tc: data=%v, k=%d, exp %d got %d", tc.data, tc.k, exp, got)
		}
	}

}
