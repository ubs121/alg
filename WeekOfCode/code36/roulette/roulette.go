// https://www.hackerrank.com/contests/w36/challenges/revised-russian-roulette

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	reader := bufio.NewReader(os.Stdin)
	buf, _ := reader.ReadBytes('\n') // read all buffer at once

	var min, max int

	l := len(buf)
	if buf[l-1] == '\n' {
		l-- // decrease by 1
	}

	i := 0

	// skip spaces
	for i < l && ' ' == buf[i] {
		i++
	}

	for i < l {
		// skip zeros (unlocked doors)
		for i < l && '0' == buf[i] {
			i++

			// skip spaces
			for i < l && ' ' == buf[i] {
				i++
			}
		}

		locked := 0
		for i < l && '1' == buf[i] {
			locked++
			i++

			// skip spaces
			for i < l && ' ' == buf[i] {
				i++
			}
		}

		// unlock operations needed
		min += locked/2 + locked%2
		max += locked
	}

	fmt.Printf("%d %d\n", min, max)
}
