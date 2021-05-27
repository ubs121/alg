package main

import (
	"strings"
	"testing"
)

type digit struct {
	all  string
	curr int
}

func newDigit(d byte) *digit {
	digitMap := map[byte]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	do := &digit{all: digitMap[d]}
	return do
}

func (d *digit) isLast() bool {
	return d.curr == len(d.all)-1
}

func (d *digit) inc() {
	d.curr++
}

func (d *digit) reset() {
	d.curr = 0
}

func (d *digit) current() byte {
	return d.all[d.curr]
}

func toString(ds []*digit) string {
	var sb strings.Builder
	for i := 0; i < len(ds); i++ {
		sb.WriteByte(ds[i].current())
	}
	return sb.String()
}

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var ret []string
	n := len(digits)

	// initial combo
	cmb := make([]*digit, n)
	for i := 0; i < n; i++ {
		cmb[i] = newDigit(digits[i] - '0')
	}

	ret = append(ret, toString(cmb)) // add initial combo

	for {
		// generate next variant
		i := 0
		for i < n && cmb[i].isLast() {
			cmb[i].reset()
			i++
		}

		if i == n {
			break
		}

		// plus one
		cmb[i].inc()

		// add into return list
		ret = append(ret, toString(cmb))
	}
	return ret
}

func TestLetterCombinations(t *testing.T) {
	testCases := map[string][]string{
		"23": {"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		"2":  {"a", "b", "c"},
		"":   {},
	}

	for tc, exp := range testCases {
		got := letterCombinations(tc)
		if !cmpStringArray(exp, got) {
			t.Errorf("tc %s: exp %v, got %v", tc, exp, got)
		}
	}
}
