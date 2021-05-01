package main

import (
	"strings"
	"testing"
)

func TestMatcher(t *testing.T) {
	testCases := map[string]bool{
		"aa a":                     false,
		"cb ?a":                    false,
		"mississippi m??*ss*?i*pi": false,
		"abcabczzzde *abc???de*":   true,
		" ?":                       false, // empty string, pattern '?'
		" ****":                    true,  // pattern '****'
		" ":                        true,  // both empty
		"a ":                       false,
		"b ?*?":                    false,
		"baba b*?a*":               true,
		"bbbbab *a?*b":             false,
		"ab *ab":                   true,
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, " ")
		got := isMatch(splits[0], splits[1])

		if got != exp {
			t.Errorf("%s: exp %v, got: %v", tc, exp, got)
		}
	}
}
