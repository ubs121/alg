package main

import "testing"

func TestSolve(t *testing.T) {
	testCases := map[int]int{
		5:     45,
		7:     27,
		1:     18,
		33:    333,
		12121: 121212,
		111:   1116,
		101:   1017,
		9:     90,
		67:    567,
		27:    207,
		80:    180,
	}

	for tc, exp := range testCases {
		got := solve(tc)
		if got != exp {
			t.Errorf("tc %d: exp %d, got %d", tc, exp, got)
		}
	}
}
