package main

import (
	"fmt"
	"index/suffixarray"
)

//https://codingcompetitions.withgoogle.com/kickstart/round/00000000008f4a94/0000000000b55464
//
// Spell: start+middle+end, each contain one or more syllables
// start==end
// start contains at least two syllables
// A syllable consists of any number of letters, including EXACTLY ONE vowel.
// There are 5 vowels: 'a', 'e', 'i', 'o' and 'u'. All other letters are considered to be consonants, including the letter 'y'

var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func solve(s string) bool {

	index := suffixarray.New([]byte(s)) // create an index for 's'
	vowelsCnt := countVowels(s)         // vowelsCnt[i] - number of vowles in range [0:i]
	dups := make(map[string]bool)       // to avoid duplicates

	// TODO: put all possible start words in hash map
	for i := 0; i < len(s)-2; i++ {
		// form a start word at 'i', at least two syllables
		j := i
		v := 0
		for j < len(s) && v < 2 {
			if vowels[s[j]] {
				v++
			}
			j++
		}
		start := s[i:j]

		if _, exists := dups[start]; exists {
			continue // skip it
		}

		// TODO: use hash map to find end words
		// find possible end words
		offsets := index.Lookup([]byte(start), -1)
		if len(offsets) > 1 {
			for _, off := range offsets { // end := s[off : off+len(start)]
				if j < off {
					// check 'middle' word
					diff := vowelsCnt[off-1] - vowelsCnt[j-1]
					if diff > 0 {
						return true // this is a spell
					}
				}
			}
		}

		dups[start] = true // mark it checked
	}

	return false
}

func countVowels(s string) []int {
	cnt := make([]int, len(s))
	prev := 0
	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			cnt[i] = prev + 1
		} else {
			cnt[i] = prev
		}
		prev = cnt[i]
	}
	return cnt
}

func main() {
	var T int
	fmt.Scanf("%d", &T)

	var word string
	for caseNumber := 1; caseNumber <= T; caseNumber++ {
		fmt.Scanf("%s", &word)
		output := solve(word)
		if output {
			fmt.Printf("Case #%d: Spell!\n", caseNumber)
		} else {
			fmt.Printf("Case #%d: Nothing.\n", caseNumber)
		}
	}
}
