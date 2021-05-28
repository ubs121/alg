package main

import (
	"strings"
	"testing"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(a, b []int) float64 {
	total := len(a) + len(b) // total elements
	rem := 0                 // number of elements to be removed from both side
	if total%2 == 0 {
		rem = (total - 2) / 2
	} else {
		rem = (total - 1) / 2
	}

	l := 0 // removed elements on the left
	r := 0 // removed elements on the right

	// safe left cut
	cutLeft := func(arr []int, n int) []int {
		if len(arr) < n {
			n = len(arr) // adjust n
		}
		if rem-l < n {
			n = rem - l // adjust with 'l'
		}
		l += n
		return arr[n:]
	}

	cutRight := func(arr []int, n int) []int {
		if len(arr) < n {
			n = len(arr)
		}
		if rem-r < n {
			n = rem - r
		}
		r += n
		return arr[:len(arr)-n]
	}

	// remove numbers from two sides of 'a' and 'b'
	// until finds the medians
	for l < rem || r < rem {
		n := len(a)
		m := len(b)
		//fmt.Printf("l=%v, r=%v, a=%v, b=%v\n", l, r, a, b)

		if m == 0 {
			if l < rem {
				a = cutLeft(a, rem-l)
			}

			if r < rem {
				a = cutRight(a, rem-r)
			}
			break
		}

		if n == 0 {
			if l < rem {
				b = cutLeft(b, rem-l)
			}

			if r < rem {
				b = cutRight(b, rem-r)
			}
			break
		}

		// Compare middle elements 'aa' and 'bb', and cut accordingly
		if a[n/2] <= b[m/2] {
			if l <= r {
				if n == 1 {
					if 1 < m {
						if a[0] <= b[m/2-1] {
							a = cutLeft(a, 1) // (a|bb*)
						} else {
							b = cutLeft(b, m/2) // (*b|ab*)
						}
					} else {
						a = cutLeft(a, 1) // (a|b)
					}
				} else if m == 1 {
					a = cutLeft(a, rem-l) // (*a|ab)
				} else {
					if b[m/2-1] <= a[n/2-1] {
						b = cutLeft(b, m/2) // (*b|aab*)
					} else {
						a = cutLeft(a, n/2) // (*a|bab, a|abb*)
					}
				}
			} else {
				if m == 1 {
					b = cutRight(b, 1) // (*aa|b)
				} else {
					b = cutRight(b, len(b)/2) // (*aa|b*)
				}
			}
		} else {
			if l <= r {
				if n == 1 {
					b = cutLeft(b, rem-l) // (*b|ba)
				} else if m == 1 {
					if 1 < n {
						if b[0] <= a[n/2-1] {
							b = cutLeft(b, 1) // (b|aa*)
						} else {
							a = cutLeft(a, n/2) // (*a|ba*)
						}
					} else {
						b = cutLeft(b, 1)
					}
				} else {
					if a[n/2-1] <= b[m/2-1] {
						a = cutLeft(a, n/2) // (*a|bba*)
					} else {
						b = cutLeft(b, m/2) // (*b|aba*, *b|baa*)
					}
				}
			} else {
				if n == 1 {
					a = cutRight(a, 1) // (*bb|a)
				} else {
					a = cutRight(a, len(a)/2) // (*bb|a*)
				}
			}
		}

		if len(a) == n && len(b) == m {
			break
		}
	}

	//fmt.Printf("Final l=%v, r=%v, a=%v, b=%v\n", l, r, a, b)

	// calculate the final answer
	sum := 0
	cnt := 0
	for _, v := range append(a, b...) {
		sum += v
		cnt++
	}

	return float64(sum) / float64(cnt)
}

func TestSolve(t *testing.T) {
	testCases := map[string]float64{
		"1,3 2":                    2.0,
		"1,2 3,4":                  2.5,
		"0,0 0,0":                  0.0,
		" 1":                       1.0,
		"2 ":                       2.0,
		" 1,2,3,4":                 2.5,
		"1,3,3,4,4 1,1,2,2,3":      2.5,
		"1 1,2,2,2,2,3,4":          2.0,
		"0,0,0,0,0 -1,0,0,0,0,0,1": 0.0,
		"1,3 2,7":                  2.5,
		"2,3,4 1":                  2.5,
		"5,6 1,2,3,4,7,8,9":        5.0,
		"1,2 -1,3":                 1.5,
		"1,4 2,3,5,6,7,8":          4.5,
		"1,3,5 2,4,6,7,8":          4.5,
		"5,6,7 1,2,3,4,8,9,10":     5.5,
		"1,2,2 1,2,3":              2.0,
		"1,2,3 1,2,2":              2.0,
	}

	for tc, exp := range testCases {
		arrays := strings.Split(tc, " ")
		a := parseIntArray(arrays[0])
		b := parseIntArray(arrays[1])

		got := findMedianSortedArrays(a, b)
		if got != exp {
			t.Errorf("%s: exp: %f, got %f", tc, exp, got)
		}
	}
}
