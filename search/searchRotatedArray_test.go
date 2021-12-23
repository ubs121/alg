package search

import (
	"alg/container"
	"strconv"
	"strings"
	"testing"
)

// https://leetcode.com/problems/search-in-rotated-sorted-array/
func searchInRotatedArray(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		mVal := nums[m]
		if mVal == target {
			return m
		}

		rVal := nums[r]
		if rVal == target {
			return r
		}

		lVal := nums[l]
		if lVal == target {
			return l
		}

		if mVal < rVal { // right is in correct order
			if mVal < target && target < rVal {
				l = m + 1 // right
			} else {
				r = m - 1 // left
			}
		} else { // left is in correct order
			if lVal < target && target < mVal {
				r = m - 1 // left
			} else {
				l = m + 1 // right
			}
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	testCases := map[string]int{
		"4,5,6,7,0,1,2|0":     4,  // right
		"4,5,6,7,8,9,0,1,2|0": 6,  // right
		"4,5,6,7,0,1,2|3":     -1, // no target
		"1|0":                 -1, // no target
		"1|1":                 0,  // middle
		"8,9,0,1,2,4,5,6,7|9": 1,  // left
		"8,9,0,1,2,4,5,6,7|0": 2,  // left
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, "|")
		arr := container.ParseIntArray(splits[0])
		target, _ := strconv.Atoi(splits[1])
		got := searchInRotatedArray(arr, int(target))
		if got != exp {
			t.Errorf("tc %s: exp %d, got %d", tc, exp, got)
		}
	}
}
