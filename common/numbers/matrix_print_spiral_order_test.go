package numbers

import (
	"reflect"
	"testing"
)

// Given an m x n matrix, return all elements of the matrix in spiral order.
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	var result []int
	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse bottom row (if applicable)
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse left column (if applicable)
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		name        string
		inputMatrix [][]int
		exp         []int
	}{
		{
			"test1",
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"test2",
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			"empty",
			[][]int{},
			[]int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := spiralOrder(tc.inputMatrix)
			if !reflect.DeepEqual(got, tc.exp) {
				t.Errorf("exp %v, but got %v", tc.exp, got)
			}
		})

	}
}
