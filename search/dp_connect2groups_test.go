package search

import (
	"math"
	"testing"
)

// https://leetcode.com/problems/minimum-cost-to-connect-two-groups-of-points/
// cost[i][j] is the cost of connecting point i of the first group and point j of the second group
func connectTwoGroups(cost [][]int) int {
	n, m := len(cost), len(cost[0])

	// memoization to avoid redundant calculations
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 1<<m) // 2^m
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// 'mask' - points in the second group that have been connected so far.
	var dfs func(int, int) int
	dfs = func(i, mask int) int {

		if dp[i][mask] != -1 {
			return dp[i][mask] // already calculated
		}

		if i < n {
			minCost := math.MaxInt32
			// Connect the current point in group 1 to some point in group 2
			for j := 0; j < m; j++ {
				// (mask|1<<j) is to set the j-th bit of the mask variable to 1.
				minCost = min(minCost, cost[i][j]+dfs(i+1, mask|(1<<j)))
			}

			dp[i][mask] = minCost
			return dp[i][mask]
		} else {
			res := 0
			for j := 0; j < m; j++ {
				if mask&(1<<j) == 0 {
					minCost := math.MaxInt32
					for k := 0; k < n; k++ {
						minCost = min(minCost, cost[k][j])
					}
					res += minCost
				}
			}
			dp[i][mask] = res
			return res
		}
	}

	return dfs(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestConnectTwoGroups(t *testing.T) {
	testCases := []struct {
		cost [][]int
		exp  int
	}{
		{
			[][]int{{15, 96}, {36, 2}},
			17,
		},
		{
			[][]int{{1, 3, 5}, {4, 1, 1}, {1, 5, 3}},
			4,
		},
		{
			[][]int{{2, 5, 1}, {3, 4, 7}, {8, 1, 2}, {6, 2, 4}, {3, 8, 8}},
			10,
		},
	}

	for i, tc := range testCases {
		got := connectTwoGroups(tc.cost)
		if tc.exp != got {
			t.Errorf("tc %d: exp %d, got %d", i, tc.exp, got)
		}
	}
}
