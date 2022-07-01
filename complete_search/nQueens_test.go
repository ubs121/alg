package search

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/n-queens/
func solveNQueens(n int) [][]string {
	out := make(chan []int)
	go func() {
		defer close(out)
		state := make([]int, n) // row -> column
		for i := 0; i < n; i++ {
			state[i] = -1 // marks it free
		}
		placeQueen(0, state, n, out)
	}()

	// empty row (all '.')
	emptyRow := make([]byte, n)
	for i := 0; i < n; i++ {
		emptyRow[i] = '.'
	}

	var ret [][]string
	for sol := range out {
		// convert into 'Q', '.' format
		var solFmt []string
		for r := 0; r < n; r++ {
			row := make([]byte, n)
			copy(row, emptyRow)
			row[sol[r]] = 'Q'
			solFmt = append(solFmt, string(row))
		}
		// collect it
		ret = append(ret, solFmt)
	}

	return ret
}

func placeQueen(r int, state []int, n int, out chan []int) {
	if r == n {
		// output
		cpy := make([]int, n)
		copy(cpy, state)
		out <- cpy
		return
	}

	// find occupied cells in row 'r'
	occCells := make([]int, n)

	for row := 0; row < r; row++ {
		// check column
		col := state[row]
		if 0 <= col {
			occCells[col] = 1
		}

		// check diagonally
		dist := r - row

		if 0 <= col-dist {
			occCells[col-dist] = 1
		}

		if col+dist < n {
			occCells[col+dist] = 1
		}
	}

	// try to place it on each free cell
	for c, occ := range occCells {
		if occ == 0 {
			state[r] = c
			placeQueen(r+1, state, n, out)
			state[r] = -1 // reset it
		}
	}
}

func TestNQueens(t *testing.T) {
	ret := solveNQueens(8)
	fmt.Printf("%v", ret)
}
