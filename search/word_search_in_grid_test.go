package search

import (
	"fmt"
	"testing"
)

// Given a 2D board and a word, determine if the word exists in the board. The word can be constructed from letters of sequentially adjacent cells.
func wordExists(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	type cell [2]int
	var search func(i, j, index int, visited map[cell]bool) bool

	search = func(i, j, index int, visited map[cell]bool) bool {
		if index == len(word) {
			return true
		}
		if i < 0 || i >= rows || j < 0 || j >= cols || visited[cell{i, j}] || board[i][j] != word[index] {
			return false
		}

		visited[cell{i, j}] = true
		result := search(i-1, j, index+1, visited) ||
			search(i+1, j, index+1, visited) ||
			search(i, j-1, index+1, visited) ||
			search(i, j+1, index+1, visited)
		delete(visited, cell{i, j})
		return result
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if search(i, j, 0, map[cell]bool{}) {
				return true
			}
		}
	}
	return false
}

func TestWordExists(t *testing.T) {
	testCases := []struct {
		grid [][]byte
		word string
		exp  bool
	}{
		{
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'}},
			word: "ABCCED",
			exp:  true,
		},
		{
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'}},
			word: "SEE",
			exp:  true,
		},
		{
			grid: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'}},
			word: "ABCB",
			exp:  false,
		},
		{
			grid: [][]byte{{'a', 'a'}},
			word: "aaa",
			exp:  false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := wordExists(tc.grid, tc.word)
			if got != tc.exp {
				t.Errorf("test case %d: exp %v got %v", i, tc.exp, got)
			}
		})
	}
}
