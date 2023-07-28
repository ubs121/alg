package search

import (
	"testing"
)

// Given an array of integers and an integer k, find the total number of continuous subarrays whose sum equals k.
func subarraySum(nums []int, k int) int {
	count := 0 // answer, number of subarrays with sum equals k
	sum := 0
	sumMap := make(map[int]int) // cumulative sum of subarrays, the cumulative sum corresponds to a distinct subarray starting from the beginning of the array
	sumMap[0] = 1               // Initialize with 0 sum, this is to account for the possibility of having a subarray with a sum of K at the very beginning of the array

	for _, num := range nums {
		// Update the sum by adding the current element.
		sum += num

		// check if sum - k exists. If it does, it means there is a subarray with sum equals k.
		if freq, exists := sumMap[sum-k]; exists {
			count += freq
		}

		// increment the frequency of 'sum'
		sumMap[sum]++
	}

	return count
}

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		output int
	}{
		{nums: []int{1, 1, 1}, k: 2, output: 2},
		{nums: []int{1, 2, 3}, k: 3, output: 2},
		{nums: []int{1, -1, 0}, k: 0, output: 3},
		{nums: []int{-1, -1, 1}, k: 0, output: 1},
		{nums: []int{}, k: 1, output: 0},
		{nums: []int{1, 2, 3, 4, 5}, k: 10, output: 1},
		{nums: []int{1, 2, 3, 4, 5}, k: 100, output: 0},
		{nums: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7, output: 4},
		{nums: []int{3, 4, 7, 2, -3, 1, 4, 2, 1}, k: 7, output: 6},
		{nums: []int{3, 0, 0, 3, 0, 0}, k: 3, output: 12},
	}

	for i, tc := range testCases {
		result := subarraySum(tc.nums, tc.k)
		if result != tc.output {
			t.Errorf("Test case %d: Expected %d, but got %d", i+1, tc.output, result)
		}
	}
}
