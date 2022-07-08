package geometric

import (
	alg "alg/container"
	"strings"
	"testing"
)

// Calculates largest rectangle under histogram
func maxRectUnderHistogram(heights []int) int {
	i, max := 0, 0
	var stack []int // left index stack

	updateMaxArea := func(rightIx int) {
		leftIx := stack[len(stack)-1] // left index
		stack = stack[:len(stack)-1]  // pop
		height := heights[leftIx]     // height
		var width int                 // width
		if len(stack) == 0 {
			width = rightIx
		} else {
			width = rightIx - stack[len(stack)-1] - 1
		}

		///<image url="maxRectUnderHistogram1.svg"/>
		// area = left.height * (rightIx - left.index - 1)
		area := height * width // current area
		if max < area {
			max = area
		}
	}

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			// height is increasing, so push it into stack
			stack = append(stack, i)
			i++
		} else {
			// update max area with respect to the top element in the stack
			updateMaxArea(i)
		}
	}

	for 0 < len(stack) {
		updateMaxArea(i) // remaining
	}
	return max
}

func TestMaxRectUnderHistogram(t *testing.T) {
	testCases := map[string]int{
		"1,3,2,2,3,0": 8,
		"1,2,3,4,5,6": 12,
	}
	for tc, exp := range testCases {
		arr := alg.ParseIntArray(tc)
		got := maxRectUnderHistogram(arr)
		if got != exp {
			t.Errorf("tc %s: exp %d, got %d", tc, exp, got)
		}
	}
}

// https://leetcode.com/problems/maximal-rectangle/
func maximalRectangle(matrix [][]int) int {
	// check if empty matrix
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	n := len(matrix)    // number of rows
	m := len(matrix[0]) // number of columns

	// temporary matrix to store histograms, or we could use original array
	tempMatrix := make([][]int, n)

	// copy 1st row and can skip
	tempMatrix[0] = make([]int, m)
	copy(tempMatrix[0], matrix[0])

	// calculate histogram for each rows
	for i := 1; i < n; i++ {
		tempMatrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				tempMatrix[i][j] = tempMatrix[i-1][j] + 1
			}
		}
	}

	// calculate maximum area from each histogram
	maxArea := 0
	for i := 0; i < n; i++ {
		max := maxRectUnderHistogram(tempMatrix[i])
		if maxArea < max {
			maxArea = max
		}
	}
	return maxArea
}

func TestMaxRectInMatrix(t *testing.T) {
	testCases := map[string]int{
		"10100\n10111\n11111\n10010":     6,
		"101101\n101110\n111110\n101101": 8,
	}
	for tc, exp := range testCases {
		matrix := parseMatrix(tc)
		got := maximalRectangle(matrix)
		if got != exp {
			t.Errorf("%s: exp %d, got %d", tc, exp, got)
		}
	}
}

// test helper that parses 2d matrix from a string
func parseMatrix(strMatrix string) [][]int {
	if len(strMatrix) == 0 {
		return nil
	}

	lines := strings.Split(strMatrix, "\n")
	matrix := make([][]int, len(lines))
	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			if line[j] == '1' {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
