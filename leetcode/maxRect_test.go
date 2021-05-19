package main

import (
	"strings"
	"testing"
)

// https://leetcode.com/problems/maximal-rectangle/

func maximalRectangle(matrix [][]byte) int {
	maxArea := 0

	// (x1, y1) is upper left corner
	for x1 := 0; x1 < len(matrix); x1++ {
		col := matrix[x1]
		for y1 := 0; y1 < len(col); y1++ {
			if col[y1] == 0 {
				continue // skip
			}
			// (x2, y2) is bottom right corner

			// TODO: find nearest non-zero point on 'x' axis (?, y1)
			// TODO: find nearest non-zero point on 'y' axis (x1, ?)
			// TODO: decrease x2, y2 by 1 and
			// but don't check if (x1,y1,x2,y2) < maxArea

			for x2 := maxArea / (y1 + 1); x2 < len(matrix); x2++ {
				for y2 := maxArea / (x1 + 1); y2 < len(col); y2++ {
					// TODO: check if 0 is in the middle

					// TODO: compare with maxArea
					area := (x2 - x1 + 1) * (y2 - y1 + 1)
					if maxArea < area {
						maxArea = area
					}

				}
			}

		}
	}
	return maxArea
}

func TestMaxRect(t *testing.T) {
	testCases := map[string]int{
		"10100\n10111\n11111\n10010\n": 6,
	}
	for tc, exp := range testCases {
		matrix := parseMatrix(tc)
		got := maximalRectangle(matrix)
		if got != exp {
			t.Errorf("%s: exp %d, got %d", tc, exp, got)
		}
	}
}

func parseMatrix(strMatrix string) [][]byte {
	if len(strMatrix) == 0 {
		return nil
	}
	splits := strings.Split(strMatrix, "\n")
	var matrix [][]byte
	for _, s := range splits {
		var row []byte
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		matrix = append(matrix, row)
	}
	return matrix
}
