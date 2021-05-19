package main

import (
	"strconv"
	"strings"
	"testing"
)

// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	// Area = length of shorter vertical line * distance between lines
	l := 0
	r := len(height)
	max := 0
	lh := 0
	rh := 0
	area := 0

	for l < r {
		lh = height[l]
		rh = height[r-1]

		if lh < rh {
			area = lh * (r - l - 1)
			l++
		} else {
			area = rh * (r - l - 1)
			r--
		}

		if max < area {
			max = area
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	testCases := map[string]int{
		"1,8,6,2,5,4,8,3,7": 49,
		"1,1":               1,
		"4,3,2,1,4":         16,
		"1,2,1":             2,
		"7,1,3,20,1,4":      21,
		"1,3,2,5,25,24,5":   24,
	}

	for tc, exp := range testCases {
		arr := parseArray(tc)
		got := maxArea(arr)
		if got != exp {
			t.Errorf("%s: exp: %d, got %d", tc, exp, got)
		}
	}
}

func parseArray(strArr string) []int {
	items := strings.Split(strArr, ",")
	var arr []int
	for i := 0; i < len(items); i++ {
		if len(items[i]) > 0 {
			n, _ := strconv.Atoi(items[i])
			arr = append(arr, n)
		}
	}
	return arr
}
