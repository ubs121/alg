package main

import "fmt"

type palindromeInRange struct {
	rangeCnt [][]int // rangeCnt[i] is number of letters in range [0:i]
}

// creates a new solver
func newPalindromeChecker(s string) *palindromeInRange {
	pc := new(palindromeInRange)
	pc.rangeCnt = make([][]int, len(s))

	r0 := make([]int, 26)
	r0[s[0]-'A']++
	pc.rangeCnt[0] = r0 // letter counts in range [0:0]

	// pre-calculate all [0:i] ranges
	for i := 1; i < len(s); i++ {
		c := s[i] - 'A'            // new character
		ri := make([]int, 26)      // new counts
		copy(ri, pc.rangeCnt[i-1]) // copy previous counts
		ri[c]++                    // increase counter
		pc.rangeCnt[i] = ri        // set it
	}

	return pc
}

// query tells if palindrome can be formed from the letters in range [l, r]
func (pc *palindromeInRange) query(l, r int) bool {
	if l == r {
		return true // single letter
	}

	if l == 1 { // all [0:r] range
		return isPalindrome(pc.rangeCnt[r-1])
	}

	lCnt := pc.rangeCnt[l-2]
	rCnt := pc.rangeCnt[r-1]

	// count letters in range [l, r]
	diff := make([]int, 26)
	for k := 0; k < 26; k++ {
		diff[k] = rCnt[k] - lCnt[k]
	}
	//fmt.Printf("l=%d, r=%d: arr=%v\n", l, r, diff)

	return isPalindrome(diff)
}

func isPalindrome(azCount []int) bool {
	odds := 0
	for _, c := range azCount {
		if c%2 == 1 {
			odds++
			// palindrome can have one or zero odd, others must be even
			if odds > 1 {
				return false
			}
		}
	}
	return true
}

// run all queries
func (pc *palindromeInRange) run(queries [][]int) int {
	n := 0
	// TODO: memorize answered queries?
	for _, q := range queries {
		if pc.query(q[0], q[1]) {
			n++
		}
	}
	return n
}

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for caseNumber := 1; caseNumber <= T; caseNumber++ {
		var N, Q int
		fmt.Scanf("%d %d", &N, &Q)

		var word string
		fmt.Scanf("%s", &word)

		var queries [][]int
		var l, r int
		for i := 0; i < Q; i++ {
			fmt.Scanf("%d %d", &l, &r)
			queries = append(queries, []int{l, r})
		}

		// create a solver and run
		pc := newPalindromeChecker(word)
		ans := pc.run(queries)
		fmt.Printf("Case #%d: %d\n", caseNumber, ans)
	}
}
