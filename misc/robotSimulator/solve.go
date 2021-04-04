package main

/*
Problem:
 A command is given in a sequence of the following instructions: a robot can go forward 'G', turn left 'L' and turn right 'R'.
 Check if given command is cycle (or repeated).
*/

// isCycle checks if cmd is cyclic
func isCycle(cmd string) bool {
	return _isCycle(0, 0, 'N', cmd, 0)
}

// (x, y) current position,
// dir - head direction
// rpt - repeats
func _isCycle(x, y int, head byte, cmd string, rpt int) bool {
	// execute all commands
	for i := 0; i < len(cmd); i++ {
		switch cmd[i] {
		case 'G': // go forward
			switch head {
			case 'N':
				y++
			case 'E':
				x++
			case 'S':
				y--
			case 'W':
				x--
			}
		case 'R': // turn right
			switch head {
			case 'N':
				head = 'E'
			case 'E':
				head = 'S'
			case 'S':
				head = 'W'
			case 'W':
				head = 'N'
			}
		case 'L': // turn left
			switch head {
			case 'N':
				head = 'W'
			case 'W':
				head = 'S'
			case 'S':
				head = 'E'
			case 'E':
				head = 'N'
			}
		}
	}

	if x == 0 && y == 0 {
		return true
	}

	// try 4 times maximum
	if rpt < 4 {
		if _isCycle(x, y, head, cmd, rpt+1) {
			return true
		}
	}

	return false
}
