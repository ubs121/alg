package numbers

import (
	"math/big"
	"strings"
	"testing"
)

func nearestPalindromic(s string) string {
	n := len(s)
	mid := n/2 + n%2    // +n%2 for odd length
	leftHalf := s[:mid] // left half of 's'

	// direct mirror
	di := reverse(leftHalf)
	di = leftHalf + di[(n%2):]

	// leftHalf-1 palindrome
	lo := strAdd(leftHalf, -1)
	if len(leftHalf) > len(lo) || lo == "0" {
		lo = strings.Repeat("9", n-1) // case 100 => 99
	} else {
		loRev := reverse(lo)
		lo = lo + loRev[(n%2):]
	}

	// leftHalf+1 palindrome
	hi := strAdd(leftHalf, 1)
	if len(leftHalf) < len(hi) {
		hi = "1" + strings.Repeat("0", n-1) + "1" // case 99 => 100
	} else {
		hiRev := reverse(hi)
		hi = hi + hiRev[(n%2):]
	}

	//fmt.Printf("s=%s, di,lo,hi=%s,%s,%s\n", s, di, lo, hi)

	if di == s { // don't include itself
		return nearestStr(s, lo, hi)
	}

	return nearestStr(s, di, lo, hi)
}

func reverse(s string) string {
	n := len(s)
	s2 := make([]byte, n) // new array for reverse
	for i := 0; i < n; i++ {
		s2[n-i-1] = s[i]
	}
	return string(s2)
}

func strAdd(s string, val int) string {
	s2 := new(big.Int)
	s2.SetString(s, 10)
	s2.Add(s2, big.NewInt(int64(val)))
	return s2.String()
}

// find closest integer (in string representation) to 's'
func nearestStr(s string, strList ...string) string {
	sBig := new(big.Int)
	sBig.SetString(s, 10)

	sMin := new(big.Int)
	sMin.SetString(strList[0], 10) // set first element as minimum

	diffMin := new(big.Int)
	diffMin = diffMin.Abs(diffMin.Sub(sMin, sBig))

	for i := 1; i < len(strList); i++ {
		bi := new(big.Int)
		bi.SetString(strList[i], 10)

		diffBi := new(big.Int)
		diffBi.Set(diffBi.Abs(diffBi.Sub(bi, sBig)))

		cmpDiff := diffMin.Cmp(diffBi)
		if cmpDiff > 0 {
			diffMin.Set(diffBi)
			sMin.Set(bi)
		} else if cmpDiff == 0 { // when diff is same
			if sMin.Cmp(bi) > 0 {
				sMin.Set(bi) // take min val
			}
		}
	}
	return sMin.String()
}

func TestClosestPalindrome(t *testing.T) {
	testCases := map[string]string{
		"123":   "121",
		"1":     "0",
		"10":    "9",
		"1000":  "999",
		"99800": "99799",
		"999":   "1001",
		"1234":  "1221",
		"12932": "12921",
		"12120": "12121",
		"99005": "98989",
	}

	for tc, exp := range testCases {
		got := nearestPalindromic(tc)

		if got != exp {
			t.Errorf("tc '%s': exp '%s', got '%s'", tc, exp, got)
		}
	}
}
