package main

import "fmt"

func solveNaive(arr []int, max int) int {
	c := 0
	for i := 0; i < len(arr)-1; i++ {

		max := arr[i] // let's take arr[i] as  max

		for j := i + 1; j < len(arr); j++ {

			// find max in segment a[i:j]
			if max < arr[j] {
				max = arr[j]
			}

			// check if a[i]*a[j]<=max
			if arr[i]*arr[j] <= max {
				c++
				fmt.Printf("%d,%d\n", i, j)
			}
		}
	}
	return c
}
