package main

import (
	"strings"
	"testing"
)

// https://leetcode.com/problems/minimum-window-substring/
// Could you find an algorithm that runs in O(n) time?
func minWindow(s string, t string) string {
	// count valid chars
	validChars := map[byte]int{}
	for i := 0; i < len(t); i++ {
		validChars[t[i]]++
	}

	seg := map[byte]int{} // window
	l := 0                // window left
	r := 0                // window right
	minStr := ""          // min window
	min := len(s)         // min window length

	// find first valid window on the left (any length)
	solFound := false
	for r < len(s) {
		if _, valid := validChars[s[r]]; valid {
			seg[s[r]]++ // count it

			// check if it's a valid window
			if len(seg) >= len(validChars) {
				if cmpWindow(seg, validChars) {
					// cut invalid or over chars from the left
					for validChars[s[l]] == 0 || seg[s[l]] > validChars[s[l]] {

						if validChars[s[l]] > 0 {
							seg[s[l]]--
						}

						l++
					}

					min = r - l + 1
					minStr = s[l : r+1]
					solFound = true
					r++
					break // break the loop
				}
			}
		}

		r++
	}

	if !solFound {
		return "" // no solution
	}

	// go over the rest and adjust the window from left and right
	for r < len(s) {

		// skip all non-valid chars
		for r < len(s) && validChars[s[r]] == 0 {
			r++
		}

		// for valid chars
		for r < len(s) && validChars[s[r]] > 0 {
			seg[s[r]]++ // count it

			if len(seg) >= len(validChars) {

				// cut invalid or over chars from the left
				for validChars[s[l]] == 0 || seg[s[l]] > validChars[s[l]] {

					if validChars[s[l]] > 0 {
						// cut from left
						c := seg[s[l]]
						c--
						if c == 0 {
							delete(seg, s[l])
						} else {
							seg[s[l]] = c
						}
					}

					l++
				}

				if r-l+1 < min {
					min = r - l + 1
					minStr = s[l : r+1]
				}
			}

			r++
		}
	}

	return minStr
}

func cmpWindow(got, exp map[byte]int) bool {
	if len(got) != len(exp) {
		return false
	}
	for k, v := range got {
		if exp[k] > v {
			return false
		}
	}
	return true
}

func TestMinWindow(t *testing.T) {
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
