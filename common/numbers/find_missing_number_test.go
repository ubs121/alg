package numbers

import (
	"alg/common"
	"testing"
)

// Given an array of integers from 1 to n with one number missing, find the missing number.
func findMissingNumber(nums []int) int {
	n := len(nums) + 1 // inlcuding a missing one

	// Calculate the expected sum of numbers from 1 to n
	expectedSum := (n * (n + 1)) / 2

	// Calculate the actual sum of the given numbers
	actualSum := common.Sum(nums)

	// The missing number is the difference between the expected sum and the actual sum
	missingNumber := expectedSum - actualSum

	return missingNumber
}

func TestFindMissingNumber(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{nums: []int{4, 1, 2}, expected: 3},
		{nums: []int{9, 6, 4, 2, 3, 5, 7, 1}, expected: 8},
		{nums: []int{1, 2, 3, 4, 6, 7, 8, 9, 10}, expected: 5},
		{nums: []int{2}, expected: 1},
		{nums: []int{4, 2, 1}, expected: 3},
	}

	for i, tc := range testCases {
		result := findMissingNumber(tc.nums)
		if result != tc.expected {
			t.Errorf("tc %d: exp %d, got %d ", i+1, tc.expected, result)
		}
	}
}
