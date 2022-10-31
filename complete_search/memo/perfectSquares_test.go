package search

import (
	"testing"
)

// https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {
	// solution store, memo[i] - minimum number of squares that sum to i
	memo := []int{0, 1}

	for len(memo) <= n {
		// TODO: find min of all prev results that are (-i*i) indices away
		memo = append(memo, 0)
	}

	return memo[n]
}

func TestNumSquares(t *testing.T) {
	testCases := map[int]int{
		12: 3,
		13: 2,
	}

	for tc, exp := range testCases {
		got := numSquares(tc)
		if exp != got {
			t.Errorf("tc [%d]: exp %v, got %v", tc, exp, got)
		}
	}
}
