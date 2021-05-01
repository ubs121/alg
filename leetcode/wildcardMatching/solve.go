package main

import (
	"index/suffixarray"
	"strings"
)

// https://leetcode.com/problems/wildcard-matching/
func isMatch(s string, p string) bool {

	// tokenize 'p'
	tokens := tokenize(p)

	// token => matches
	matches := make([][]int, len(tokens))

	// a suffix array for 's'
	index := suffixarray.New([]byte(s))

	mPos := 0 // last matching position

	// match all tokens in sequence
	for i, tkn := range tokens {
		if tkn[0] == '*' {
			continue // skip
		}

		// match at the begining
		if i == 0 {
			if !hasPrefix(s, tkn) {
				return false // no match
			}
			matches[i] = []int{0}
			continue
		}

		// match at the end
		if i == len(tokens)-1 {
			if !hasPrefix(s[len(s)-len(tkn):], tkn) {
				return false // no match
			}
			matches[i] = []int{len(s) - len(tkn)}
			continue
		}

		// remove prefix ? marks for prefix search
		preQM := 0
		for i := 0; i < len(tkn) && tkn[i] == '?'; i++ {
			preQM++
		}
		pttn1 := tkn[preQM:] // cut ? marks

		// all question marks ?
		if len(pttn1) == 0 {
			// TODO:
			continue
		}

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
			if mPos+preQM <= ix && mPos+ix+len(tkn) <= len(s) {
				// validate if actually matches
				if hasPrefix(s[ix:], tkn[preQM:]) {
					ixs1 = append(ixs1, ix)
				}
			}
		}

		if len(ixs1) == 0 {
			return false // no match actually after cleanup
		}

		matches[i] = ixs1

	}

	if mPos < len(s) {
		// TODO: check if s ends with '*'
	}

	return true
}

// tokenize pattern
func tokenize(p string) []string {
	var tokens []string
	i := 0
	for i < len(p) {
		if p[i] == '*' {
			// skip all '*' sequence
			for i < len(p) && p[i] == '*' {
				i++
			}
			tokens = append(tokens, "*") // add as one '*' token
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
