package main

import "math"
import "fmt"

// TODO: Kadane algorithm

// Find the maximum possible sum in arr[] such that arr[m] is part of it
func maxCrossingSum(arr []int, l, m, h int) int {
	// Include elements on left of mid.
	sum := 0
	leftSum := math.MinInt32
	for i := m; i >= l; i-- {
		sum = sum + arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}


	// Include elements on right of mid
	sum = 0
	rightSum := math.MinInt32
	for i := m + 1; i <= h; i++ {
		sum = sum + arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	// Return sum of elements on left and right of mid
	return leftSum + rightSum
}

// Returns sum of maxium sum subarray in arr[l..h]
func maxSubArraySum(arr []int, l, h int) int {
	// Base Case: Only one element
	if l == h {
		return arr[l]
	}

	// Find middle point
	m := (l + h) / 2

	/* Return maximum of following three possible cases
	   a) Maximum subarray sum in left half
	   b) Maximum subarray sum in right half
	   c) Maximum subarray sum such that the subarray crosses the midpoint */
	return max(maxSubArraySum(arr, l, m),
		maxSubArraySum(arr, m+1, h),
		maxCrossingSum(arr, l, m, h))
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(a, b, c int) int {
	return max2(max2(a, b), c)
}

/*Driver program to test maxSubArraySum*/
func main() {
	arr := []int{1, 1, -20, 1, 20}
	max_sum := maxSubArraySum(arr, 0, len(arr)-1)
	fmt.Printf("Maximum contiguous sum is %d\n", max_sum)
}
