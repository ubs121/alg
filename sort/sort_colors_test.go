package sort

import (
	"fmt"
	"testing"
)

// Sort Colors: Given an array with n objects colored red, white, or blue,
// sort them in-place so that objects of the same color are adjacent,
// with the colors in the order red, white, and blue.
func sortColors(nums []int) {
	count := make([]int, 3) // Count the occurrences of each color

	// Count the number of occurrences of each color
	for _, color := range nums {
		count[color]++
	}

	index := 0

	// Fill the nums array with the sorted colors
	for color := 0; color <= 2; color++ {
		for count[color] > 0 {
			nums[index] = color
			index++
			count[color]-- // descrease the counter
		}
	}
}

func TestSortColors(t *testing.T) {
	colors := []int{0, 2, 2, 1, 2, 0, 1}
	sortColors(colors)
	fmt.Printf("colors=%v", colors)
}
