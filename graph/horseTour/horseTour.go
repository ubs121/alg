package main

var (
	board [8][8]byte
	done  bool
)

func main() {
	move(0, 0, 1)
}

func move(x, y int, n byte) {
	if !done && 0 <= x && x < 8 && 0 <= y && y < 8 && board[x][y] == 0 {
		board[x][y] = n
		if n == 64 {
			printBoard()
			done = true
			return
		}
		move(x-1, y+2, n+1)
		move(x-2, y+1, n+1)
		move(x-2, y-1, n+1)
		move(x-1, y-2, n+1)

		move(x+1, y-2, n+1)
		move(x+2, y-1, n+1)
		move(x+2, y+1, n+1)
		move(x+1, y+2, n+1)

		board[x][y] = 0
	}
}

func printBoard() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			print(" ", board[i][j])
		}
		println()
	}
}
