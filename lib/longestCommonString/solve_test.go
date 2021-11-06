package main

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {

	testCases := map[string]int{
		"HARRY,SALLY":       2,
		"AA,BB":             0,
		"SHINCHAN,NOHARAAA": 3,
		"HHNAN,NHAAA":       2,
		"AAABBB,BBBAAA":     3,
		"CAAAAAB,BAAAAAC":   5,
	}

	for in, exp := range testCases {
		inputs := strings.Split(in, ",")
		got := commonChild(inputs[0], inputs[1])
		if got != exp {
			t.Errorf("[%s] expected %d, but got %d", in, exp, got)
		}
	}
}
