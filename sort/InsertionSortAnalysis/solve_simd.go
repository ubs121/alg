package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	var n int

	for t := 0; t < T; t++ {
		fmt.Scanf("%d\n", &n)

		a := make([]int, n) // input array

		lim := 0
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &a[i])
			if lim < a[i] {
				lim = a[i]
			}
		}
		lim++

		c := make([]int, lim) // counter array

		s := 0 // total shifts
		for i := 0; i < n; i++ {
			s += c[a[i]] // shifts

			// TODO: use SIMD or broadcast !!!
			for j := 0; j < a[i]; j++ {
				c[j]++
			}
		}

		fmt.Printf("%d\n", s)
	}
}
