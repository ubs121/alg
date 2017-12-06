// https://www.hackerrank.com/contests/w28/challenges/lucky-number-eight
package main

import "fmt"

var c int
var n int
var a []int

/**
*	i - digit position
* 	x - previuos value
* 	t - which tenth
 */
func check(i, x, t int) {
	x = a[i]*t + x // concate a number at position 'i'

	if x%8 == 0 {
		c++
		//fmt.Println(x)
	}

	for ii := i + 1; ii < n; ii++ {
		check(ii, x, t*10)
	}
}

func solve(x int) int {

	a = nil

	for x > 0 {
		a = append(a, x%10)
		x = x / 10
	}

	c = 0
	n = len(a)

	for i := 0; i < n; i++ {

		// check only even numbers
		if a[i]%2 == 0 {
			check(i, 0, 1)
		}
	}

	return c
}

func solve2(m int) int {
	a = nil

	for m > 0 {
		a = append(a, m%10)
		m = m / 10
	}

	c = 0
	n = len(a)

	var j, p, d, x int

	for i := 1; i < 64; i++ {
		//fmt.Printf("%06b\n", i)

		j = i
		p = 1
		d = 0
		x = 0 // generated number
		for j > 0 && d < n {
			// a[d] цифр орно
			if j%2 == 1 {

				x = a[d]*p + x

				if p == 1 && a[d]%2 == 1 {
					// last digit is odd
					break
				}

				p = p * 10

			}

			d++
			j = j / 2
		}

		//  check generated number
		if j == 0 && x%8 == 0 {
			c++
			//fmt.Println(x)
		}

	}

	return c
}

func solve3(m int) int {
	a = nil

	for m > 0 {
		a = append(a, m%10)
		m = m / 10
	}

	c = 0
	n = len(a)

	for 
}

func main() {
	t := 2 * 100000
	for i := 1; i < t; i++ {
		s1 := solve(i)
		s2 := solve2(i)

		if s1 != s2 {
			fmt.Println(i, "->", s1, s2)
		}
	}
}
