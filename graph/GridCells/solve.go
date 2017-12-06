// https://www.hackerrank.com/challenges/connected-cell-in-a-grid
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int // dimension (n * m)

	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &m)

	// read array
	scanner := bufio.NewScanner(os.Stdin)
	a := make(map[int]int)

	for i := 0; i < n && scanner.Scan(); i++ { // should be n lines
		vals := strings.Split(scanner.Text(), " ")
		for j := 0; j < len(vals); j++ {
			v, _ := strconv.Atoi(vals[j])
			if v > 0 {
				a[i*m+j] = 1 // i and j are connected
			}
		}
	}

	var track func(x, y int) int
	track = func(x, y int) int {
		c := 0

		if x >= 0 && y >= 0 && x < n && y < m {
			k := x*m + y
			if _, exists := a[k]; exists {
				delete(a, k) // mark it visited

				// count it
				c = 1
				c += track(x-1, y-1)
				c += track(x-1, y)
				c += track(x-1, y+1)
				c += track(x, y-1)
				c += track(x, y+1)
				c += track(x+1, y-1)
				c += track(x+1, y)
				c += track(x+1, y+1)
			}
		}
		return c
	}

	max := 1
	x := 0
	y := 0
	for len(a) > max { // while the map is not empty
		//fmt.Printf("%v\n", a)

		// find starting cell (hack)
		for k := range a {
			x = k / m
			y = k % m

			break
		}

		c := track(x, y)
		if c > max {
			max = c
		}
	}

	fmt.Println(max)
}
