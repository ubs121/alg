package main

import (
	"testing"
)

func TestPalindromeChecker(t *testing.T) {
	type testcase struct {
		input   string
		queires [][]int
	}
	testCases := map[*testcase]int{
		{"ABAACCA", [][]int{{3, 6}, {4, 4}, {2, 5}, {6, 7}, {3, 7}}}:                         3,
		{"XYZ", [][]int{{1, 3}, {1, 3}, {1, 3}, {1, 3}, {1, 3}}}:                             0,
		{"XADGALAGDAX", [][]int{{1, 11}, {2, 10}, {3, 9}, {1, 3}, {1, 1}, {11, 11}, {2, 3}}}: 5,
	}
	for tc, exp := range testCases {
		pc := newPalindromeChecker(tc.input)
		got := pc.run(tc.queires)
		if got != exp {
			t.Errorf("tc %s: exp %d, got %d", tc.input, exp, got)
		}
	}
}

// func main2() {
// 	var T int
// 	fmt.Scanf("%d", &T)
// 	for caseNumber := 1; caseNumber <= T; caseNumber++ {
// 		// TODO: read input
// 		var arr []byte
// 		pc := newPalindromeChecker(arr)

// 		// TODO: read quires
// 		var queries [][]int
// 		ans := pc.run(queries)
// 		fmt.Printf("Case #%d: %d\n", caseNumber, ans)
// 	}
// }
