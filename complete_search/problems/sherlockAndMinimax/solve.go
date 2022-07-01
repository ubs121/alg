// https://www.hackerrank.com/challenges/sherlock-and-minimax

package main

import "fmt"
import "sort"

func sub(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func min(x int, a []int) int {
	m := sub(a[0], x)
	d := m
	for i := 1; i < len(a); i++ {
		d = sub(a[i], x)
		if m > d {
			m = d
		}
	}

	return m
}

func main() {
	var n int // n >= 1
	var p, q int

	fmt.Scanf("%d", &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	fmt.Scanf("%d %d", &p, &q)

	sort.Ints(a) // sort data

	var m int // M number

	if q <= a[0] {
		m = p
	} else if p >= a[n-1] {
		m = q
	} else {
		var d, _m int

		m = p
		pmin := min(p, a)
		qmin := min(q, a)

		mmin := pmin
		if pmin < qmin {
			m = q
			mmin = qmin
		}

		for i := 1; i < n; i++ {
			_m = (a[i]+a[i-1])/2 + (a[i]+a[i-1])%2
			if p <= _m && _m <= q {
				d = sub(a[i], _m)

				if mmin < d {
					m = _m
					mmin = d

					//fmt.Println(m, mmin)
				}
			}
		}
	}

	fmt.Println(m)

}
