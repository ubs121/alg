// Frog jumping
package main

import (
	"fmt"
	"testing"
)

const (
	river string = "RWWRRRWWRWWRRRRRRRRRRRRRRWWW"
	vel   int    = 2
)

var (
	mem map[int]bool
)

func frogJump(pos int) bool {

	if pos >= len(river) {
		return true
	}

	if pos >= 0 && river[pos] == 'W' {
		return false
	}

	// loop over [vel-1, vel, vel+1]
	for v := vel - 1; v <= vel+1; v++ {
		if b, ok := mem[pos+v]; ok { // check from memory
			return b
		} else {
			mem[pos+v] = frogJump(pos + v) // calculate

			if mem[pos+v] {
				return true
			}
		}
	}

	return false
}

func TestFrogJump(t *testing.T) {

	mem = make(map[int]bool)

	if frogJump(-1) { // starting position
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
