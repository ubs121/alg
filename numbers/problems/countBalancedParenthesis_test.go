package numbers

import (
	"alg/common"
	"fmt"
	"testing"
)

// https://codingcompetitions.withgoogle.com/kickstart/round/00000000008f4a94/0000000000b5496b
func countValidParenthesis(l, r int) int {
	// odd brakets divides the whole string, so don't put in the middle
	// no nesting because it divides
	m := common.Min(l, r)
	sol := m * (m + 1) / 2
	return sol
}

func main() {
	var T int
	fmt.Scanf("%d", &T)

	var l, r int
	for caseNumber := 1; caseNumber <= T; caseNumber++ {
		fmt.Scanf("%d %d", &l, &r)
		output := countValidParenthesis(l, r)
		fmt.Printf("Case #%d: %d\n", caseNumber, output)
	}
}

func TestCoundValidParenthesis(t *testing.T) {
	testCases := map[string]int{
		"1,0": 0,
		"1,1": 1,
		"2,2": 3,
		"3,2": 3,
		"5,7": 15,
	}

	for tc, exp := range testCases {
		inpArr := common.ParseIntArray(tc)
		got := countValidParenthesis(inpArr[0], inpArr[1])
		if got != exp {
			t.Errorf("tc [%s]: exp %d, got %d", tc, exp, got)
		}
	}
}
