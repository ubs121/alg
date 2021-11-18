package main

import (
	"fmt"
	"strings"
	"testing"
)

// https://leetcode.com/problems/wildcard-matching/
func isMatch(s string, p string) bool {
	i, j := 0, 0 // indices to be matched in s and p
	star := -1
	star_i := -1

	for i < len(s) {
		if j >= len(p) || (p[j] != '*' && p[j] != '?' && p[j] != s[i]) {
			if star == -1 {
				return false
			}
			j = star + 1
			star_i++
			i = star_i
		} else if p[j] == '*' {
			star = j
			star_i = i
			j++
		} else {
			i++
			j++
		}
	}

	for j < len(p) {
		if p[j] != '*' {
			return false
		}
		j++
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
