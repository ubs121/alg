package main

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
