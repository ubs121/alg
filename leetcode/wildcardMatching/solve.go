package main

import (
	"index/suffixarray"
	"strings"
)

// https://leetcode.com/problems/wildcard-matching/
func isMatch(s string, p string) bool {

	// a suffix array for 's'
	index := suffixarray.New([]byte(s))

	// tokenize 'p'
	tokens := tokenize(p)

	// pattern => match indexes, matches[i] == nil if tokens[i]=='*'
	matches := make([][]int, len(tokens))

	// collect all matches for each non '*' token
	for i, pttn := range tokens {
		if pttn[0] != '*' {

			// match at the begining
			if i == 0 {
				if !hasPrefix(s, pttn) {
					return false // no match
				}
				matches[i] = []int{0}
				continue
			}

			// match at the end
			if i == len(tokens)-1 {
				if !hasPrefix(s[len(s)-len(pttn):], pttn) {
					return false // no match
				}
				matches[i] = []int{len(s) - len(pttn)}
				continue
			}

			// remove prefix ? marks for prefix search
			preQMark := 0
			for i := 0; i < len(pttn) && pttn[i] == '?'; i++ {
				preQMark++
			}
			pttn1 := pttn[preQMark:] // cut ? marks

			// still has a question mark?
			hasQM := strings.IndexByte(pttn1, '?')
			if hasQM > 0 {
				pttn1 = pttn1[:hasQM] // cut it
			}

			ixs := index.Lookup([]byte(pttn1), -1)
			if ixs == nil {
				return false
			}

			// cleanup 'ixs' checking with whole 'pttn'
			var ixs1 []int
			for _, ix := range ixs {
				if preQMark <= ix && ix+len(pttn) < len(s) {
					// validate if actually matches
					if hasPrefix(s[ix:], pttn[preQMark:]) {
						ixs1 = append(ixs1, ix)
					}
				}
			}

			if len(ixs1) == 0 {
				return false // no match actually after cleanup
			}

			matches[i] = ixs1
		}

	}

	// find a correct arrangement (a path) for whole pattern
	matchFound := pathExists(matches, 0, -1)
	if matchFound {
		// must verify if last pattern is covering till the end
		lastInd := len(tokens) - 1
		lastToken := tokens[lastInd]

		if lastToken == "*" {
			return true // it will stretch to the end
		}

		ixs := matches[lastInd]
		for i := 0; i < len(ixs); i++ {
			if ixs[i]+len(lastToken) == len(s) {
				return true
			}
		}
		return false
	}

	return false
}

// check if there is a path from beginning to end in 'matches'
func pathExists(matches [][]int, level, prevIndex int) bool {
	// reached at last level ?
	if level == len(matches)-1 {
		return true
	}

	ixs := matches[level]
	if len(ixs) == 0 {
		// skip, this is '*' node
		return pathExists(matches, level+1, prevIndex)
	}

	for _, ix := range matches[level] {
		if prevIndex < ix { // take greater indexes only
			if pathExists(matches, level+1, ix) {
				return true
			}
		}
	}

	return false
}

// tokenize pattern
func tokenize(p string) []string {
	var tokens []string
	i := 0
	for i < len(p) {
		if p[i] == '*' {
			tokens = append(tokens, "*") // '*' token
			i++
		} else {
			j := i + 1
			for j < len(p) && p[j] != '*' {
				j++
			}
			tokens = append(tokens, p[i:j]) // letters, ?
			i = j
		}
	}
	return tokens
}

// checks if 's' has prefix 'p', pattern supports ? only
func hasPrefix(s string, p string) bool {
	if len(p) > len(s) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if s[i] != p[i] && p[i] != '?' {
			return false
		}
	}
	return true
}
