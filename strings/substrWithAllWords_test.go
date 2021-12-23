package strings

import (
	"strings"
	"testing"
)

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {

	wLen := len(words[0])       // length of one word
	segLen := wLen * len(words) // length of search window

	validWords := map[string]int{}
	for _, w := range words {
		validWords[w]++
	}

	var ans []int // solution

	for i := 0; i < len(s)-segLen+1; i++ {

		seg := s[i : i+segLen]     // current segment
		segCnt := map[string]int{} // word counts in the segment
		countOk := true            // true if word numbers are equal or less than expected
		for j := 0; j < segLen-wLen+1; j += wLen {
			w := seg[j : j+wLen]
			segCnt[w]++
			if c, valid := validWords[w]; !valid || segCnt[w] > c {
				countOk = false
				break
			}
		}

		if countOk && len(segCnt) == len(validWords) {
			ans = append(ans, i) // match found
		}
	}
	return ans
}

func TestFindSubstring(t *testing.T) {
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
