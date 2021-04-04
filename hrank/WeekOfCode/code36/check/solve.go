// https://www.hackerrank.com/contests/w36/challenges/ways-to-give-a-check

package main

import (
	"bufio"
	"fmt"
	"os"
)

const n = 8

type piece struct {
	kind byte // piece type
	x, y int  // position
}

func ways2check(board [][]byte) int {

	var blackKing, pawn piece
	var all []piece

	// find the pawn
	for j := 0; j < n; j++ {
		if board[1][j] == 'P' && board[0][j] == '#' {
			pawn = piece{kind: 'P', x: 0, y: j} // suppposed position
			board[1][j] = '#'                   // remove
			break
		}
	}

	// find the black king and others
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c := board[i][j]
			if c != '#' {
				p := piece{kind: c, x: i, y: j}
				if c == 'k' {
					blackKing = p
				} else {
					all = append(all, p) // except the pawn & black king
				}
			}
		}
	}

	check := func(p1, p2 piece) bool {
		dx := p1.x - p2.x
		dy := p1.y - p2.y

		if dx != 0 {
			dx = dx / abs(dx)
		}

		if dy != 0 {
			dy = dy / abs(dy)
		}

		x := p2.x + dx
		y := p2.y + dy

		for board[x][y] == '#' {
			x += dx
			y += dy
		}

		return (board[x][y] == p1.kind)
	}

	// look for "discovered checks" from white pieces
	for _, p := range all {
		if 'a' <= p.kind && p.kind <= 'z' {
			continue // black piece
		}

		// white queen or rook
		if p.kind == 'Q' || p.kind == 'R' {
			if blackKing.x == p.x || blackKing.y == p.y {
				if check(blackKing, p) {
					return 4
				}
			}
		}

		// white queen or bishop
		if p.kind == 'Q' || p.kind == 'B' {
			if abs(p.x-blackKing.x) == abs(p.y-blackKing.y) {
				if check(blackKing, p) {
					return 4
				}
			}
		}
	}

	// check attack from promoted pawn
	dx := pawn.x - blackKing.x
	dy := pawn.y - blackKing.y

	ans := 0

	// 1. Pawn to Knight
	if (abs(dx) == 1 && abs(dy) == 2) || (abs(dx) == 2 && abs(dy) == 1) {
		ans++
	}

	// 2. Pawn to Rook or Queen
	if dx == 0 || dy == 0 {
		if check(blackKing, pawn) {
			ans += 2
		}
	}

	// 3. Pawn to Bishop
	if abs(dx) == abs(dy) {
		if check(blackKing, pawn) {
			ans += 2
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var t int
	fmt.Scanf("%d\n", &t)

	reader := bufio.NewReader(os.Stdin)

	for q := 0; q < t; q++ {

		board := make([][]byte, n)
		for i := 0; i < n; i++ {
			buf, _ := reader.ReadBytes('\n')
			board[i] = buf
		}

		// check ways
		fmt.Println(ways2check(board))
	}

}
