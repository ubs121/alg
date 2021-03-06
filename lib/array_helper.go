package alg

import "math"

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
	for i := 0; i < len(arr)/2; i++ {
		arr[len(arr)-i-1], arr[i] = arr[i], arr[len(arr)-i-1]
	}
}
