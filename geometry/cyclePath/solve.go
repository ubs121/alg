package main

import "math"

/*
Problem:
 A command is given in a sequence of the following instructions: a robot can go forward 'G', turn left 'L' and turn right 'R'.
 Check if given command is cycle (or repeated).
*/

const (
	North = 0
	West  = 1
	South = 2
	East  = 3
)

// isCycle checks if cmd is cyclic
func isCycle(cmd string) bool {
	x, y := 0, 0 // origin
	dir := North // direction

	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case 'R': // turn right
			switch dir {
			case 'N':
				dir = East
			case 'E':
				dir = South
			case 'S':
				dir = West
			case 'W':
				dir = North
			}
		case 'L': // turn left
			switch dir {
			case 'N':
				dir = West
			case 'W':
				dir = South
			case 'S':
				dir = East
			case 'E':
				dir = North
			}
		case 'G': // go forward
			switch dir {
			case 'N':
				y++
			case 'E':
				x++
			case 'S':
				y--
			case 'W':
				x--
			}
		}
	}

	if x == 0 && y == 0 && dir == North {
		return true // came to the origin, and with same direction, so this is a cycle
	}

	if dir == South {
		return true // if apply this path one more time it will come back to the origin
	}

	l := math.Sqrt(float64(x*x + y*x))

	// apply this path 4 times at maximum, it should find if there's a cycle
	for i := 0; i < 4; i++ {
		switch dir {
		case South:
			y = int(y - l)
		}
	}

	return false
}
