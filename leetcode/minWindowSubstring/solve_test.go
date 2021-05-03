package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := map[string]string{
		"ADOBECODEBANC ABC": "BANC",
		"a a":               "a",
		"a aa":              "",
		"ab b":              "b",
		"bba ab":            "ba",
		"abc b":             "b",
		"AbabbbAbaA Bab":    "",
	}
	for tc, exp := range testCases {
		splits := strings.Split(tc, " ")
		got := minWindow(splits[0], splits[1])
		if got != exp {
			t.Errorf("%s: exp %s, got %s", tc, exp, got)
		}
	}
}
