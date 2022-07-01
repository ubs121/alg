package geometric

import (
	"testing"
)

// https://leetcode.com/problems/robot-bounded-in-circle/
func isRobotBounded(cmd string) bool {
	x, y := 0, 0 // position
	head := 'N'  // head direction

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

	// after all moves
	// it's a cycle if the robot is at origin or direction is the same
	return (x == 0 && y == 0) || head != 'N'
}

func TestIsRobotBounded(t *testing.T) {
	commands := map[string]bool{
		"GLL":    true,
		"RG":     true,
		"GR":     true,
		"RGGRGG": true, // will be cycle after 8 repeats
		"GGGGG":  false,
		"LLGRL":  true,
	}

	for tc, exp := range commands {
		got := isRobotBounded(tc)
		if got != exp {
			t.Errorf("cmd=%s: exp %v, got %v", tc, exp, got)
		}
	}
}
