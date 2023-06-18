package numbers

import (
	alg "alg/common"
	"fmt"
	"math"
	"math/bits"
	"strings"
	"testing"
)

func TestBits(t *testing.T) {
	total := uint64(0)
	sum, carry := uint(0), uint(0)
	digits := 3 // # of digits
	//maxDigits:=[]int{} // actual digits of n
	for i := 0; i < int(math.Pow(2, float64(digits)))-1; i++ {
		sum, carry = bits.Add(sum, 1, carry)
		zeros := bits.LeadingZeros(sum) - (64 - digits)
		wild := digits - bits.OnesCount(sum)
		if wild > 0 {
			if zeros > 0 {
				// TODO: not 8, up to maxDigit[x]
				total += 8 * uint64(math.Pow(9, float64(wild-1)))
			} else {
				total += uint64((math.Pow(9, float64(wild))))
			}
		} else {
			total++ // all 1s
		}

		fmt.Printf("sum = %03b, zeros=%d, *=%d, total=%d\n", sum, zeros, wild, total)
	}
	fmt.Printf("total=%d", total)
}

func countDigitOne(n int) int {
	return 0
}

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
		if !alg.CmpUnorderedStringArray(exp, got) {
			t.Errorf("tc %s: exp %v, got %v", tc, exp, got)
		}
	}
}
