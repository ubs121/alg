package alg

import (
	"testing"
)

var dictIntToRoman = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var dictRomanToInt = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

// https://leetcode.com/problems/integer-to-roman/
// 1 <= num <= 3999
func intToRoman(num int) string {
	// Roman numerals are usually written largest to smallest from left to right.
	// I can be placed before V (5) and X (10) to make 4 and 9.
	// X can be placed before L (50) and C (100) to make 40 and 90.
	// C can be placed before D (500) and M (1000) to make 400 and 900.

	ret := ""
	r := 0    // mod 10
	d := 1    // # of digits
	sym := "" // roman symbol for 'r'

	for num > 0 {
		r = num % 10

		if r == 4 || r == 9 || r == 5 {
			sym = dictIntToRoman[d*r]
			ret = sym + ret
		} else {
			if 5 < r {
				tmp := dictIntToRoman[d*5]
				r = r - 5
				sym = dictIntToRoman[d]
				for i := 0; i < r; i++ {
					tmp = tmp + sym
				}
				ret = tmp + ret
			} else {
				sym = dictIntToRoman[d]
				for i := 0; i < r; i++ {
					ret = sym + ret
				}
			}
		}

		num = num / 10
		d = d * 10
	}

	return ret
}

func romanToInt(s string) int {
	num := 0
	i := len(s)
	for i > 1 {
		if n, exists := dictRomanToInt[s[i-2:i]]; exists {
			num += n
			i--
		} else {
			num += dictRomanToInt[s[i-1:i]]
		}
		i--
	}

	// last digit
	if i > 0 {
		num += dictRomanToInt[s[i-1:i]]
	}

	return num
}

func TestInt2Roman(t *testing.T) {
	testCases := map[int]string{
		3:    "III",
		4:    "IV",
		9:    "IX",
		58:   "LVIII",
		1994: "MCMXCIV",
	}

	for tc, exp := range testCases {
		got := intToRoman(tc)
		if got != exp {
			t.Errorf("tc %d: exp %s, got %s", tc, exp, got)
		}
	}

	for exp, tc := range testCases {
		got := romanToInt(tc)
		if got != exp {
			t.Errorf("tc %s: exp %d, got %d", tc, exp, got)
		}
	}
}
