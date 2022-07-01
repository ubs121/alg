package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestIrregularExpressions(t *testing.T) {
	testCases := map[string]bool{
		"abracadabra":       true,
		"kajabbamajabbajab": true,
		"frufrumfuffle":     false,
		"schprexityschprex": false,
		"aauaa":             true,
		"aaaa":              false,
	}

	for tc, exp := range testCases {
		got := solve(tc)
		if got != exp {
			t.Errorf("tc %s: exp %v, got %v", tc, exp, got)
		}
	}
}

func TestStringSplit(t *testing.T) {
	s := "kajabbamajabbajab"
	tokens := strings.FieldsFunc(s, func(c rune) bool {
		return vowels[byte(c)]
	})
	fmt.Printf("tokens=%v", tokens)
}
