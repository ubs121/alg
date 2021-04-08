package main

import "testing"

func TestSolve(t *testing.T) {
	testCases := [][]int{
		{1, 1, 5},
		{2, 2, 3, 7},
		{10, 7, 12},
		{2, 5, 5, 5, 5, 5},
		{5, 5, 5, 5, 5},
	}
	exp := []int{2, 2, 3, 6, 0}

	for i := 0; i < len(testCases); i++ {
		got := equal(testCases[i])
		if exp[i] != got {
			t.Errorf("%v: exp %d, got %d", testCases[i], exp[i], got)
		}
	}
}
