package main

import "fmt"

func main() {
	var n, c int
	var a []int
	var s string

	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)

	// inverse digits into an array
	a = make([]int, len(s))
	for i := 0; i < n; i++ {
		a[i] = int(s[n-i-1] - '0')
	}

	var j, p, d, x int
	for i := 1; i < 128; i++ {

		j = i
		p = 1
		d = 0
		x = 0 // generated number
		for j > 0 && d < n {
			// a[p] цифр орно
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
			fmt.Printf("%d -> %06b\n", x, i)
		}

	}

	fmt.Println(c)
}
