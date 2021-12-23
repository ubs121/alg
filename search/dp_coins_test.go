// https://www.hackerrank.com/challenges/coin-change
// https://www.hackerrank.com/challenges/ctci-coin-change/problem

package search

import (
	"testing"
)

var (
	// solution store ( amount -> ( first_k_coins -> number_of_variants) )
	mem map[int]map[int]int
)

func coinChange(n int, coins []int) int {

	// check from the solution store
	if arrSol, ok := mem[n]; ok && arrSol != nil {
		if k, ok := arrSol[len(coins)]; ok {
			return k
		}
	}

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	k := 0

	// бага дугаартай зоосноос эхлэн задлах
	for i := 0; i < len(coins); i++ {
		k = k + coinChange(n-coins[i], coins[i:]) // remove duplicates c[i:]
	}

	// remember in the solution store
	if _, exists := mem[n]; !exists {
		mem[n] = make(map[int]int)
	}

	mem[n][len(coins)] = k

	return k
}

// iterative constructing approach
func coinChange2(n int, coins []int) int {
	count := make([]int, n+1) // solution store
	count[0] = 1
	for x := 1; x <= n; x++ {
		for _, c := range coins {
			if x-c >= 0 {
				count[x] += count[x-c]
			}
		}
	}
	return count[n]
}

func TestCoinChange(t *testing.T) {
	// input format
	type input struct {
		coins []int
		x     int
	}
	testCases := map[*input]int{
		{
			coins: []int{1, 3, 4},
			x:     5,
		}: 6,
	}

	for tc, exp := range testCases {
		got := coinChange2(tc.x, tc.coins)
		if exp != got {
			t.Errorf("tc: %v, exp %d, got %d", tc, exp, got)
		}
	}
}
