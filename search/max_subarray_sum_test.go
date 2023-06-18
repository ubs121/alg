package search

import (
	"math"
	"testing"
)

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray2(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum += nums[i]
		if currentSum < nums[i] {
			currentSum = nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func TestMaxSubarray(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		exp  int
	}{
		{arr: []int{1, 1, -20, 1, 20}, exp: 21},
		{arr: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, exp: 6},
		{arr: []int{5, 4, -1, 7, 8}, exp: 23},
	}
	for _, tc := range testCases {
		got := maxSubArray2(tc.arr)
		if got != tc.exp {
			t.Errorf("exp %d, got %d\n", tc.exp, got)
		}
	}
}

func maxSubArray(nums []int) int {
	return maxSubArraySum(nums, 0, len(nums)-1)
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
	return max(max(maxSubArraySum(arr, l, m), maxSubArraySum(arr, m+1, h)), maxCrossingSum(arr, l, m, h))
}

// Find the maximum possible sum in arr[] such that arr[m] is part of it
func maxCrossingSum(arr []int, l, m, h int) int {
	// elements on the left
	sum := 0
	leftSum := math.MinInt32
	for i := m; i >= l; i-- {
		sum = sum + arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	// elements on the right
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
