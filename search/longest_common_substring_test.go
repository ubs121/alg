package search

import (
	"fmt"
	"testing"
)

func longestCommonSubstring(str1, str2 string) string {
	m := len(str1)
	n := len(str2)

	// dp[i][j] represents the length of the longest common substring between str1[:i] and str2[:j].
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Variables to keep track of the longest common substring
	maxLength := 0
	endIndex := 0

	// Compute the lengths of common substrings using dynamic programming
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i - 1
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	// Extract the longest common substring from the input strings
	longestSubstring := str1[endIndex-maxLength+1 : endIndex+1]
	return longestSubstring
}

func TestLongestCommonSubstring(t *testing.T) {
	str1 := "ABAB"
	str2 := "BABA"

	longestSubstring := longestCommonSubstring(str1, str2)
	fmt.Println("Longest Common Substring:", longestSubstring)
	// Output: Longest Common Substring: ABA
}
