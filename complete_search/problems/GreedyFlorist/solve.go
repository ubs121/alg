// https://www.hackerrank.com/challenges/greedy-florist
package main

import (
	"fmt"
	"sort"
)

func multSum(vect []int, val int) int {
	s := 0
	for i := 0; i < len(vect); i++ {
		s = s + vect[i]*val
	}
	return s
}

func main() {
	var n, k int

	fmt.Scanf("%d %d", &n, &k)

	c := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &c[i])
	}

	// sort prices in descending order
	sort.Slice(c, func(i, j int) bool { return c[i] > c[j] })

	p := 0 // total price
	x := 1 // price multiplier

	for len(c) > k { // buy until last flower !!!
		p = p + multSum(c[:k], x) // calculate price
		x = x + 1
		c = c[k:]
	}

	p = p + multSum(c[:], x)

	fmt.Printf("%d\n", p)
}
