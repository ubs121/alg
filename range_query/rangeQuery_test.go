package search

import (
	"alg/container"
	"alg/numbers"
	"math"
	"strings"
	"testing"
)

// Query a sparse table in range [l,r)
func minQuery(sparseTable [][]int, l, r int) int {
	p := int(math.Log(float64(r - l + 1)))
	return numbers.Min(sparseTable[p][l], sparseTable[p][r-(1<<p)])
}

// Builds a sparse table for Range Minimum Query
func buildSparseTable(arr []int) [][]int {
	n := len(arr)                                   // # of columns
	h := int(math.Floor(math.Log2(float64(n)))) + 1 // # of rows (height)
	st := make([][]int, h)                          // sparse table, st[k][j]=min(range(j:j+2^k))

	// base case: 2^0
	st[0] = make([]int, n)
	copy(st[0], arr)

	// iterative dynamic programming approach
	for k := 1; k < h; k++ {
		st[k] = make([]int, n-(1<<k)+1) // cut to actual length
		for j := 0; j+(1<<k) <= n; j++ {
			st[k][j] = numbers.Min(st[k-1][j], st[k-1][j+(1<<(k-1))])
		}
	}
	return st
}

func TestSparseTableRMQ(t *testing.T) {
	testCases := map[string]int{
		"3,1,5,3,4,7,6,1|3,5": 3,
		"3,1,5,3,4,7,6,1|2,8": 1,
		"1|0,1":               1,
	}

	for tc, exp := range testCases {
		strArr := strings.Split(tc, "|")
		arr := container.ParseIntArray(strArr[0])
		rng := container.ParseIntArray(strArr[1])

		st := buildSparseTable(arr)
		got := minQuery(st, rng[0], rng[1])
		if exp != got {
			t.Errorf("tc: %s exp: %d got %d", tc, exp, got)
		}
	}
}
