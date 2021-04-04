package main

import "fmt"

var n int
var a []int
var c int

/**
*	i - digit position
* 	x - previuos value
* 	t - which tenth
 */
func check(i, x, t int) {
	x = a[i]*t + x // concate a number at position 'i'

	if x%8 == 0 {
		c++
		fmt.Println(x)

		if t == 100 {
			if x%8 == 0 {
				switch n - 3 {
				case 1:
					c = c + 1
				case 2:
					c = c + 3
				case 3:
					c = c + 7
				case 4:
					c = c + 15
				}
			}

			return // break
		}
	}

	for ii := i + 1; ii < n; ii++ {
		check(ii, x, t*10)
	}
}

func main() {

	var s string

	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)

	// inverse digits into an array
	a = make([]int, len(s))
	for i := 0; i < n; i++ {
		a[i] = int(s[n-i-1] - '0')
	}

	for i := 0; i < n; i++ {

		// check only even numbers
		if a[i]%2 == 0 {
			check(i, 0, 1)
		}
	}

	fmt.Println(c)

}
