package alg

import (
	"math/big"
	"strings"
	"testing"
)

// https://leetcode.com/problems/multiply-strings/
func multiplyStrings(num1 string, num2 string) string {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString(num1, 10)
	b.SetString(num2, 10)

	a.Mul(a, b)

	return a.String()
}

func TestMultiplyStrings(t *testing.T) {
	testCases := map[string]string{
		"123*456": "56088",
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, "*")
		got := multiplyStrings(splits[0], splits[1])
		if got != exp {
			t.Errorf("tc %s: exp %s, got %s", tc, exp, got)
		}
	}
}
