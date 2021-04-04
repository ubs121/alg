// https://www.hackerrank.com/contests/w35/challenges/triple-recursion
package main

import (
	"fmt"
)

func main() {
	var n, m, k int

	fmt.Scanf("%d %d %d\n", &n, &m, &k)

	a := make([]int, n)

	for i := 0; i < n; i++ {
		// 0 -> i
		j := 0
		for j < i {
			a[j]--
			fmt.Printf("%d ", a[j])
			j++
		}

		a[j] = m
		t := a[j]
		// i -> n
		for j < n {
			fmt.Printf("%d ", t)
			t--
			j++
		}

		// increase 'm'
		m = m + k
		fmt.Println()
	}
}
