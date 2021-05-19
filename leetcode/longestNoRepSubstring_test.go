package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := map[string]int{
		"eeydgwdykpv": 7,
	}
	for tc, exp := range testCases {
		got := lengthOfLongestSubstring(tc)
		if exp != got {
			t.Errorf("tc %s: exp %d, got %d, ", tc, exp, got)
		}
	}
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	maxLen := 0
	cnt := map[byte]int{}
	r := 0
	l := 0
	for {
		// increase 'r'
		r++
		for r <= len(s) {
			cnt[s[r-1]]++
			if len(cnt) == r-l {
				if maxLen < r-l {
					maxLen = r - l
				}
			} else {
				break
			}
			r++
		}

		// break if 'r' reached to the end
		if r > len(s) {
			break
		}

		// increase 'l'
		for l < len(s) && len(cnt) < r-l {
			cnt[s[l]]--
			if cnt[s[l]] == 0 {
				delete(cnt, s[l])
			}
			l++
		}
	}
	return maxLen
}
