package main

import "testing"

func TestSolve(t *testing.T) {
	testCases := map[string]int{
		"abba":       4,
		"abcd":       0,
		"cdcd":       5,
		"kkkk":       10,
		"ifailuhkqq": 3,
		"dbcfibibcheigfccacfegicigcefieeeeegcghggdheichgafhdigffgifidfbeaccadabecbdcgieaffbigffcecahafcafhcdg": 1464,
	}

	for k, exp := range testCases {
		got := sherlockAndAnagrams(k)
		if got != exp {
			t.Errorf("%s: exp %d, got %d", k, exp, got)
		}
	}
}
