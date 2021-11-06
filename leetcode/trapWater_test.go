package main

import (
	alg "alg/lib"
	"testing"
)

func trapMaxWater(height []int) int {
	// Water can be trapped inside the current and the next equal or higher bar
	// iterate till the next bar and calculate the volume in between them

	n := len(height)

	if n < 2 {
		return 0
	}

	// calculate right maxes, to avoid repeated checks
	rightMaxes := make([]int, n) // rightMaxes[i] - right max in range height[i+1:]
	maxSoFar := 0
	rightMaxes[n-1] = maxSoFar

	for i := n - 1; i > 0; i-- {
		if maxSoFar < height[i] {
			maxSoFar = height[i]
		}
		rightMaxes[i-1] = maxSoFar
	}

	curr := 0
	vol := 0
	totalVol := 0

	for i := 0; i < n; i++ {

		if curr > height[i] { // iterate till the next bar and calculate the volume in between them
			vol += (curr - height[i])
		} else {
			totalVol += vol // accumalte the volume
			vol = 0         // reset local volume

			// choose next 'curr', checking with right maxes
			j := i
			for j < n && height[j] > rightMaxes[j] {
				j++
			}

			if j < n {
				if i < j {
					curr = rightMaxes[j]
				} else {
					curr = height[i]
				}
			} else {
				break
			}

		}
	}

	totalVol += vol

	return totalVol
}

func TestTrapMaxWater(t *testing.T) {
	testCases := map[string]int{
		"4,2,3":                   1,
		"0,1,0,2,1,0,1,3,2,1,2,1": 6,
		"4,2,0,3,2,5":             9,
	}

	for tc, exp := range testCases {
		arr := alg.ParseIntArray(tc)
		got := trapMaxWater(arr)
		if got != exp {
			t.Errorf("tc %s: exp %d, got %d", tc, exp, got)
		}
	}
}
