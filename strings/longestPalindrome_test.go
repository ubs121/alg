package strings

import (
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := map[string]string{
		"babad":        "bab|aba",
		"cbbd":         "bb",
		"a":            "a",
		"bb":           "bb",
		"ccc":          "ccc",
		"aaaa":         "aaaa",
		"tattarrattat": "tattarrattat",
	}
	for tc, exp := range testCases {
		got := longestPalindrome(tc)
		expSet := map[string]bool{}
		splits := strings.Split(exp, "|")
		for _, s := range splits {
			expSet[s] = true
		}
		if _, exists := expSet[got]; !exists {
			t.Errorf("tc %s: exp %s, got %s, ", tc, exp, got)
		}
	}
}

// https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	maxLen := 1
	maxPal := s[:1]
	n := len(s)

	for i := 1; i < n; i++ {
		if 2*i < maxLen && 2*(n-i)+1 < maxLen {
			break // no need check further
		}

		// odd check, find max palindrome centered at 'i'
		maxAt := 1
		d := 1
		for i-d >= 0 && i+d < n {
			if s[i-d] != s[i+d] {
				break
			}

			maxAt += 2
			if maxLen < maxAt {
				maxLen = maxAt
				maxPal = s[i-d : i+d+1]
			}
			d++
		}

		// even check
		maxAt = 0
		d = 1
		for i-d >= 0 && i+d <= n {
			if s[i-d] != s[i+d-1] {
				break
			}

			maxAt += 2
			if maxLen < maxAt {
				maxLen = maxAt
				maxPal = s[i-d : i+d]
			}
			d++

		}

	}
	return maxPal
}
