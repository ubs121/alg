// https://www.hackerrank.com/challenges/newyear-game/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(a []int, n int) string {
	// count remainders
	r := make([]int, 3) // r0, r1, r2
	for i := 0; i < n; i++ {
		r[a[i]%3]++
	}

	fmt.Printf("%v\n", r)

	// calculate
	if r[1] > 1 {
		return "Koca"
	}

	if r[2] > 1 {
		return "Koca"
	}

	if r[1] == 0 && r[2] == 0 { // r[0] > 0
		return "Koca"
	}

	return "Balsa"
}

func main() {
	var T int
	fmt.Scanf("%d\n", &T)

	a := make([]int, 2000)
	reader := bufio.NewReader(os.Stdin)

	for t := 0; t < T; t++ {
		// n
		buf, _ := reader.ReadBytes('\n')
		n, _ := readInt(buf)

		// a[n]
		buf, _ = reader.ReadBytes('\n')
		off := 0
		for i := 0; i < n; i++ {
			a[i], off = readInt(buf)
			buf = buf[off:]
		}

		fmt.Println(solve(a, n))
	}
}

func readInt(buf []byte) (int, int) {
	off := 0
	l := len(buf)
	for off < l && ' ' == buf[off] {
		off++
	}

	d := 0 // number
	for off < l && '0' <= buf[off] && buf[off] <= '9' {
		d = d*10 + int(buf[off]-'0')
		off++
	}
	return d, off
}
