// https://www.hackerrank.com/challenges/pairs?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=7-day-campaign
package main

import (
	"fmt"
	"sort"
)

// Problem: count (i, j) pairs that a[i] - a[j] == k
func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	rem := make(map[int][]int)

	var p, r int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &p)

		// Partition by reminder (a[i] % k is partition number)
		r = p % k
		// add 'p' into a partition 'r'
		rem[r] = append(rem[r], p)
	}

	c := make(chan int) // counter channel

	// count in all partitions
	for _, v := range rem {
		//fmt.Printf("%v\n", v)
		go func(v []int, c chan int) {
			m := 0       // the number of pairs
			sort.Ints(v) // sort numbers

			for i := 1; i < len(v); i++ {
				if v[i]-v[i-1] == k {
					m++
				}
			}

			c <- m // send sum to c
		}(v, c)
	}

	s := 0
	for i := 0; i < len(rem); i++ {
		s = s + <-c
	}

	fmt.Println(s)
}
