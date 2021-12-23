package main

import "testing"

func TestSolve(t *testing.T) {
	in := [][]int{
		{4, 8, 2},
		{4, 5, 7},
		{6, 1, 6}}

	got := formingMagicSquare(in)
	exp := 4
	if got != exp {
		t.Errorf("exp %d, got %d", exp, got)
	}

	in = [][]int{
		{2, 9, 8},
		{4, 2, 7},
		{5, 6, 7}}

	got = formingMagicSquare(in)
	exp = 21
	if got != exp {
		t.Errorf("exp %d, got %d", exp, got)
	}
}
