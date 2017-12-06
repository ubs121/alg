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
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &a[i])
		}

		ch := make(chan int)

		gt := func(a []int, x int, c chan int) {
			p := 0
			for j := 0; j < len(a); j++ {
				if a[j] > x {
					p++
				}
			}
			c <- p
		}

		for i := 1; i < n; i++ {
			go gt(a[:i], a[i], ch)
		}

		s := 0
		for i := 1; i < n; i++ {
			s += <-ch
		}

		fmt.Printf("%d\n", s)
	}
}
