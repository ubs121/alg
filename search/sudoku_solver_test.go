package search

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/sudoku-solver/
func solveSudoku(board [][]byte) bool {
	return solve(board)
}

func solve(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				// If a cell is empty ('.'), try filling it with numbers from '1' to '9' and recursively calls itself to check if the number is valid and leads to a solution.
				for num := '1'; num <= '9'; num++ {
					if isValid(board, row, col, byte(num)) {
						// If the number is valid, proceed to the next empty cell and repeat the process.
						board[row][col] = byte(num)
						if solve(board) {
							return true // solution found
						}
						// If number is invalid or the board is already filled, backtrack to the previous cell and try a different number
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

// checks the validity of a number in a specific row, column
func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
		if board[row][i] == num {
			return false
		}
		// row/3 = row group, i/3 = individual cell within the group
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

func TestSudokuSolver(t *testing.T) {
	testCases := []struct {
		board [][]byte
	}{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
		},
		{
			board: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			solved := solveSudoku(tc.board)
			if !solved {
				t.Errorf("test case %d: exp 'true' got '%v'", i, solved)
			}
		})
	}
}
