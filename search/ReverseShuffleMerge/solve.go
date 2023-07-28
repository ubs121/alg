// https://www.hackerrank.com/challenges/reverse-shuffle-merge
package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)
	sb := []byte(s)
	reverse(sb[:]) // reverse it

	// disctinct letters in 's', (letter :-> count)
	dist := make(map[byte]int)
	for i := 0; i < len(sb); i++ {
		if j, exists := dist[sb[i]]; exists {
			dist[sb[i]] = j + 1
		} else {
			dist[sb[i]] = 1
		}
	}

	req := make([]byte, 0)  // required letters
	shuf := make([]int, 26) // filler letters

	// collect letters of reverse & shuffle
	for l, c := range dist {
		for i := 0; i < c/2+c%2; i++ {
			req = append(req, l)
			shuf[l-'a']++
		}
	}
	// 'req' should be ordered
	sort.Slice(req, func(i, j int) bool { return req[i] < req[j] })

	flr := make([]int, 26) // temporary filler
	ans := make([]byte, 0) // answer of problem
	next := 0              // next letter in req
	for 0 < len(req) {
		copy(flr, shuf) // copy current state of 'shuf'

		// check if 'next' is possible ?
		nextOk := true
		j := 0
		for j < len(sb) && sb[j] != req[next] {
			if flr[sb[j]-'a'] > 0 {
				flr[sb[j]-'a']--
			} else {
				// not enough filler
				nextOk = false
				break
			}
			j++
		}

		// if it's okay
		if nextOk {
			ans = append(ans, req[next])                    // collect answer
			copy(shuf, flr)                                 // set state as 'flr'
			req = req[:next+copy(req[next:], req[next+1:])] // remove 'next' from 'req'
			next = 0                                        // start from the first letter in 'req'

			if j < len(sb) {
				sb = sb[j+1:] // cut 'sb' by 'j'
			}
		} else {
			// copy(shuf, shuf) keep original state
			next++ // try next letter
		}
	}

	fmt.Printf("%s\n", ans)
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[len(s)-i-1], s[i] = s[i], s[len(s)-i-1]
	}
}
