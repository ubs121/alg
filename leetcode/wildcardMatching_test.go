package main

import (
	"fmt"
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

// https://leetcode.com/problems/regular-expression-matching/submissions/
func isMatch2(s string, p string) bool {
	tkns := tokenize2(p)
	return _match2(s, tkns, true)
}

// regular expression matching with support for '.' and '*'
func _match2(s string, tkns []string, matchAtBeginning bool) bool {
	if len(tkns) == 0 {
		return (len(s) == 0 || !matchAtBeginning)
	}

	tkn := tkns[0]
	switch tkn {
	case ".*":
		return _match2(s, tkns[1:], false)
	case ".":
		if matchAtBeginning {
			return len(s) > 0 && _match2(s[1:], tkns[1:], true)
		} else {
			// check all cuts in s[1:]
			i := 1
			for i <= len(s) {
				if _match2(s[i:], tkns[1:], true) {
					return true
				}
				i++
			}
		}
	default:
		if strings.HasSuffix(tkn, "*") { // c*
			c := tkn[0]
			i := 0
			if !matchAtBeginning {
				if _match2(s, tkns[1:], matchAtBeginning) { // c{0} case first
					return true
				}

				i = strings.IndexByte(s, c) // c+ case
				if i < 0 {
					return false
				}
				i++
			}

			// count 'c'
			for i < len(s) && s[i] == c {
				i++
			}

			// check all cuts in s[j:i]
			j := 0
			for j <= i {
				if _match2(s[j:], tkns[1:], true) {
					return true
				}
				j++
			}
		} else { // literal
			if matchAtBeginning {
				return strings.HasPrefix(s, tkn) && _match2(s[len(tkn):], tkns[1:], true)
			} else {
				// check all cuts positions of 'tkn'
				for len(s) > 0 {
					ix := strings.Index(s, tkn)
					if ix < 0 {
						return false
					}

					s = s[ix+len(tkn):]
					if _match2(s, tkns[1:], true) {
						return true
					}
				}
			}
		}
	}

	return false
}

func tokenize2(p string) []string {
	var tokens []string
	var sb strings.Builder

	for i := 0; i < len(p); i++ {
		switch p[i] {
		case '.':
			if sb.Len() > 0 {
				tokens = append(tokens, sb.String())
				sb.Reset()
			}
			if i+1 < len(p) && p[i+1] == '*' {
				tokens = append(tokens, ".*")
				i++
			} else {
				tokens = append(tokens, ".")
			}

		default:
			if i+1 < len(p) && p[i+1] == '*' {
				if sb.Len() > 0 {
					tokens = append(tokens, sb.String())
					sb.Reset()
				}
				tokens = append(tokens, p[i:i+2])
				i++
			} else {
				sb.WriteByte(p[i])
			}
		}
	}

	if sb.Len() > 0 {
		tokens = append(tokens, sb.String())
	}
	return tokens
}

func TestMatcher2(t *testing.T) {
	testCases := map[string]bool{
		"ab .*":                  true,
		"aa a":                   false,
		"aab c*a*b":              true,
		"mississippi mis*is*p*.": false,
		"abcabczzzde abc.*ee*":   true,
		"aaa a*a":                true,
		"aa a*":                  true,
		"b .*b.*":                true,
		" .":                     false,
		"a ":                     false,
		"aabc .*.c":              true,
		"aaba ab*a*c*a":          false,
		"abcdede ab.*de":         true,
		"a .*.":                  true,
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, " ")
		got := isMatch2(splits[0], splits[1])

		if got != exp {
			t.Errorf("%s: exp %v, got: %v", tc, exp, got)
		}
	}
}

func TestTokenize2(t *testing.T) {
	tkns := tokenize2("mis*is*p*.")
	fmt.Printf("%v", tkns)
}
