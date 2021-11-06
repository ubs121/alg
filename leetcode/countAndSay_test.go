// https://leetcode.com/problems/count-and-say/
package main

import (
	"strconv"
	"testing"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	// get previous string
	str := countAndSay(n - 1)
	m := len(str)

	// say it
	say := make([]byte, 0, 2*len(str))
	i := 0
	for i < m {
		ch := str[i] // current char
		c := 1       // count it
		i++          // move to next char
		for i < m && str[i] == ch {
			c++
			i++
		}

		// say it
		cStr := strconv.Itoa(c)
		say = append(say, []byte(cStr)...)
		say = append(say, ch)
	}
	return string(say)
}

func TestCountAndSay(t *testing.T) {
	testCases := map[int]string{
		4:  "1211",
		10: "13211311123113112211",
	}

	for tc, exp := range testCases {
		got := countAndSay(tc)
		if got != exp {
			t.Errorf("tc %d: exp %s, got %s", tc, exp, got)
		}
	}
}
