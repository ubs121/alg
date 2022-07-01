package strings

import (
	"strconv"
	"strings"
	"testing"
)

// hashing function for anagrams
func hashAnagram(s string) string {
	buf := make([]byte, 26) // could be []rune here
	for i := 0; i < len(s); i++ {
		buf[s[i]-'a']++
	}
	return string(buf)
}

// hashing function for anagrams
func hashAnagram2(s string) string {
	// count
	var buf [26]int
	for i := 0; i < len(s); i++ {
		buf[s[i]-'a']++
	}

	// hash
	var sb strings.Builder
	for i := 0; i < 26; i++ {
		if buf[i] > 0 {
			sb.WriteByte(byte('a' + i))
			sb.WriteString(strconv.Itoa(buf[i]))
		}
	}
	return sb.String()
}

func TestHashAnagram(t *testing.T) {
	testCases := map[string]string{
		"eat": "ate",
		"bat": "tab",
	}

	for tc, ana := range testCases {
		got := hashAnagram(tc)
		exp := hashAnagram(ana)
		if got != exp {
			t.Errorf("%s vs %s: exp true, got false", tc, ana)
		}
	}
}
