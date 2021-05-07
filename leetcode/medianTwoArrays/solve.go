package main

import (
	"sort"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(a, b []int) float64 {
	total := len(a) + len(b) // total elements
	rem := 0                 // number of elements to be removed
	if total%2 == 0 {
		rem = (total - 2) / 2
	} else {
		rem = (total - 1) / 2
	}

	l := 0 // removed elements on the left
	r := 0 // removed elements on the right

	cutLeft := func(arr []int, n int) []int {
		if len(arr) < n {
			n = len(arr) // adjust n
		}
		l += n
		return arr[n:]
	}

	cutRight := func(arr []int, n int) []int {
		if len(arr) < n {
			n = len(arr)
		}
		r += n
		return arr[:len(arr)-n]
	}

	// remove numbers from two sides of 'a' and 'b'
	for len(a) > 1 && len(b) > 1 && (l < rem || r < rem) {
		n := len(a)
		m := len(b)
		//fmt.Printf("l=%v, r=%v, a=%v, b=%v\n", l, r, a, b)

		// Compare 4 middle elements 'aa' and 'bb'
		if a[n/2] <= b[m/2] {

			if l < r {
				if b[m/2-1] <= a[n/2-1] {
					if l+m/2 <= rem {
						b = cutLeft(b, m/2) // (baab)
					}
				} else {
					if l+n/2 <= rem {
						a = cutLeft(a, n/2) // (abab, aabb)
					}
				}
			} else {

				if r+len(b)-len(b)/2 < rem {
					b = cutRight(b, len(b)/2) // *b
				}
			}
		} else {

			if l < r {
				if a[n/2-1] <= b[m/2-1] {
					if l+n/2 <= rem {
						a = cutLeft(a, n/2) // (abba)
					}
				} else {
					if l+m/2 <= rem {
						b = cutLeft(b, m/2) // (baba, bbaa)
					}
				}
			} else {

				if r+len(a)-len(a)/2 < rem {
					a = cutRight(a, len(a)/2) // *a
				}
			}
		}

		if len(a) == n && len(b) == m {
			break
		}
	}

	//fmt.Printf("Final l=%v, r=%v, a=%v, b=%v\n", l, r, a, b)

	// calculate the final answer
	a = append(a, b...)
	sort.Ints(a)

	if l < rem {
		a = cutLeft(a, rem-l)
	}

	if r < rem {
		a = cutRight(a, rem-r)
	}

	sum := 0
	cnt := 0

	for _, v := range a {
		sum += v
		cnt++
	}

	return float64(sum) / float64(cnt)
}
