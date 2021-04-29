package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := map[string][]int{
		"barfoofoobarthefoobarman bar foo the":                                      {6, 9, 12},
		"wordgoodgoodgoodbestword word good best word":                              {},
		"wordgoodgoodgoodbestword word good best good":                              {8},
		"lingmindraboofooowingdingbarrwingmonkeypoundcake fooo barr wing ding wing": {13},
		"aaaaaaaaaaaaaa aa aa":                                                      {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		"aaaaaaaa aa aa aa":                                                         {0, 1, 2},
		"ababaab ab ba ba":                                                          {1},
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, " ")
		got := findSubstring(splits[0], splits[1:])

		if !equal(got, exp) {
			t.Errorf("%s: exp %v, got: %v", tc, exp, got)
		}
	}
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
