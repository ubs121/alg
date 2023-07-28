package common

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	set2 := map[int]bool{4: true, 5: true, 6: true, 7: true, 8: true}

	intersection := FindIntersection(set1, set2)
	fmt.Println("Intersection:", intersection)
}

// We have a list of words - [abcd, abc. acdef]. find the common alphabets from the list of words
func commonAlphabets(words []string) []byte {
	if len(words) == 0 {
		return nil
	}

	set := make(map[byte]bool) // common set

	// inti the set from the first word
	for i := 0; i < len(words[0]); i++ {
		set[words[0][i]] = true
	}

	// iterate over the words
	for i := 1; i < len(words); i++ {

		// delete non-common alphabets from the set
		for k := range set {
			ix := bytes.IndexByte([]byte(words[i]), k)
			if ix < 0 {
				delete(set, k)
			}
		}
	}

	var result []byte
	for k := range set {
		result = append(result, k)
	}

	return result
}

func TestCommonAlphabets(t *testing.T) {
	testCases := []struct {
		words []string
		exp   string
	}{
		{
			words: []string{"abcd", "abc.", "acdef"},
			exp:   "ac",
		},
		{
			words: []string{"abc.", "acdef", "abcd"},
			exp:   "ac",
		},
	}
	for i, tc := range testCases {
		got := commonAlphabets(tc.words)

		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})

		if string(got) != tc.exp {
			t.Errorf("%d: exp %s, got %s", i, tc.exp, got)
		}
	}
}
