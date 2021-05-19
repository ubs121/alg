package main

import (
	"index/suffixarray"
	"sort"
	"strings"
	"testing"
)

// https://leetcode.com/problems/wildcard-matching/
func isMatch(s string, p string) bool {

	// tokenize 'p'
	tokens := tokenize(p)

	// a suffix array for 's'
	index := suffixarray.New([]byte(s))
	n := len(s) // string length
	mPos := 0   // last matching position

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
			mPos = 0 + len(tkn)
			continue
		}

		// match at the end
		if i == len(tokens)-1 {
			if mPos+len(tkn) > n {
				return false
			}

			mPos = n - len(tkn)
			if !hasPrefix(s[mPos:], tkn) {
				return false // no match
			}
			mPos = n
			continue
		}

		tkn1 := strings.TrimLeft(tkn, "?") // trim ? marks for prefix search
		nqm := len(tkn) - len(tkn1)        // # of question marks in front

		// all question mark token ?
		if len(tkn1) == 0 {
			// take first match
			mPos += len(tkn)
			if mPos > n {
				return false // can't fit
			}
			continue
		}

		// still has a question mark?
		if hasQM := strings.IndexByte(tkn1, '?'); hasQM > 0 {
			tkn1 = tkn1[:hasQM] // cut it
		}

		ixs := index.Lookup([]byte(tkn1), -1)
		if ixs == nil {
			return false
		}
		sort.Ints(ixs)

		// check 'ixs' against pattern, gready search
		ixFound := false
		for _, ix := range ixs {
			if mPos+nqm <= ix && ix+len(tkn1) <= n {
				// check if actually matches
				if hasPrefix(s[ix:], tkn[nqm:]) {
					mPos = ix + len(tkn)
					ixFound = true
					break
				}
			}
		}

		if !ixFound {
			return false // no match actually after cleanup
		}

	}

	if mPos < n {
		// check if s ends with '*'
		if !strings.HasSuffix(p, "*") {
			return false
		}
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

// https://leetcode.com/problems/regular-expression-matching/submissions/
func isMatch2(s string, p string) bool {

	tokens := strings.Split(s, ".*")
	for i, tkn := range tokens {
		// continue
		if len(tkn) == 0 {
			continue // must be prefix or suffix ".*"
		}

		if i == 0 {
			// TODO:  match at the beginning, this must be non ".*" start
			continue
		}

		// TODO: recurse for each indexOf() variants
		// TODO: break on first complete solution, otherwise return no solution
	}

	return true
}

// checks if 's' matches pattern 'p', that supports "." and "*"
func hasPrefix2(s string, p string) bool {
	i := 0
	j := 0

	ix := strings.Index(s, ".*")
	if ix >= 0 {
		s = s[:ix]
	}

	for i < len(p) {
		if p[i] == '.' {
			i++
			j++
			continue
		}

		if p[i] == '*' {
			// skip all last char
			prevChar := s[j-1]
			for j < len(s) && prevChar == s[j] {
				j++
			}
			i++
		}

		if p[i] == s[j] {
			j++
		}

		i++
	}
	return true
}

func TestMatcher1(t *testing.T) {
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

func TestMatcher2(t *testing.T) {
	testCases := map[string]bool{
		"ab .*":                  true,
		"aa a*":                  true,
		"aa a":                   false,
		"aab c*a*b":              true,
		"mississippi mis*is*p*.": false,
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, " ")
		got := isMatch2(splits[0], splits[1])

		if got != exp {
			t.Errorf("%s: exp %v, got: %v", tc, exp, got)
		}
	}
}
