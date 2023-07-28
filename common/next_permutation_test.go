package common

import (
	"reflect"
	"testing"
)

// Given a list of integers, find the next permutation in lexicographic order.
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Find the first pair where nums[i] < nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		// Find the first element greater than nums[i]
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// Swap the elements at indices i and j
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse the sublist starting from index i+1
	// This ensures that the next permutation is as small as possible while still being greater than the current permutation
	Reverse(nums[i+1:])
}

func TestNextPermutation(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			nums:     []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
		{
			nums:     []int{1},
			expected: []int{1},
		},
		{
			nums:     []int{1, 3, 2},
			expected: []int{2, 1, 3},
		},
	}

	for _, tc := range testCases {
		nextPermutation(tc.nums)
		if !reflect.DeepEqual(tc.nums, tc.expected) {
			t.Errorf("Expected %v, but got %v", tc.expected, tc.nums)
		}
	}
}
