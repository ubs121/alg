package common

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// Returns sum
func Sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

// IndexMax returns max value indexes
func IndexMax(arr []int) []int {
	var maxes []int
	maxVal := math.MinInt64 // default max

	for i := 0; i < len(arr); i++ {
		if maxVal < arr[i] {
			maxVal = arr[i] // new max
			maxes = nil     // another max found, so reset
		} else if maxVal == arr[i] {
			// collect max positions
			maxes = append(maxes, i)
		}
	}

	return maxes
}

// Reverse array elements
func Reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// Parses text into integer array
func ParseIntArray(strArr string) []int {
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

func CmpUnorderedStringArray(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
