package main

import (
	"testing"
)

func TestSolve(t *testing.T) {

	arr := []int{33, 11, 44, 11, 55}
	queries := []int{1, 2, 3, 4, 5}

	exp := []int{11, 33, 44, 44, 55}
	ret := solve(arr, queries)

	for i, got := range ret {
		if got != exp[i] {
			t.Errorf("[query %d]: expected %d, but got %d", queries[i], exp[i], got)
		}
	}

	arr = []int{1, 2, 3, 4, 5}
	queries = []int{1, 2, 3, 4, 5}

	exp = []int{1, 2, 3, 4, 5}
	ret = solve(arr, queries)

	for i, got := range ret {
		if got != exp[i] {
			t.Errorf("[query %d]: expected %d, but got %d", queries[i], exp[i], got)
		}
	}

	arr = []int{5, 4, 1, 3, 4, 1, 2, 3}
	queries = []int{1, 2, 3, 7, 8}

	exp = []int{1, 2, 3, 4, 5}
	ret = solve(arr, queries)

	for i, got := range ret {
		if got != exp[i] {
			t.Errorf("[query %d]: expected %d, but got %d", queries[i], exp[i], got)
		}
	}
}

func TestUpdateRanks(t *testing.T) {
	arr := []int{5, 2, 1, 3, 4, 1, 2, 3}
	rank := make([]int, len(arr))

	// default rank for each element
	for i := 0; i < len(rank); i++ {
		rank[i] = 1
	}

	updateRanks(rank, arr)

	exp := []int{8, 2, 1, 3, 7, 1, 2, 3}
	updateRanks(rank, arr)

	for i, got := range rank {
		if got != exp[i] {
			t.Errorf("[%d] expected %d, but got %d", arr[i], exp[i], got)
		}
	}
}
