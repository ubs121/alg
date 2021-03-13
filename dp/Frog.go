package main

import "fmt"

const (
	river string = "RWWRRRWWRWWRRRRRRRRRRRRRRWWW"
	vel   int    = 2
)

var (
	mem map[int]bool
)

func jump(pos int) bool {

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
			mem[pos+v] = jump(pos + v) // calculate

			if mem[pos+v] {
				return true
			}
		}
	}

	return false
}

func main() {

	mem = make(map[int]bool)

	if jump(-1) { // starting position
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
