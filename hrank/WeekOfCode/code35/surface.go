// https://www.hackerrank.com/contests/w35/challenges/3d-surface-area/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, w int
	fmt.Scanf("%d %d\n", &h, &w)
	reader := bufio.NewReader(os.Stdin)

	a := make([][]int, h+2)
	a[0] = make([]int, w+2)   // first row
	a[h+1] = make([]int, w+2) // last row
	// [[0 0 0 0 0]
	//  [0 1 3 4 0]
	//  [0 2 2 3 0]
	//  [0 1 2 4 0]
	//  [0 0 0 0 0]]

	ch := make(chan int)
	for i := 1; i <= h; i++ {
		a[i] = make([]int, w+2)
		buf, _ := reader.ReadBytes('\n')
		convertRow(buf, a[i], w)

		// sum up previuos row
		go sumRow(a, h, w, i-1, ch)
	}

	// last row
	go sumRow(a, h, w, h, ch)

	area := 0
	for i := 0; i <= h; i++ {
		area += <-ch
	}

	fmt.Printf("%d\n", area)
}

func sumRow(a [][]int, h, w int, i int, ch chan int) {
	sum := 0

	for j := 1; i > 0 && j <= w; j++ {
		p := 2 // top & bottom

		if a[i][j] > a[i][j-1] {
			p += a[i][j] - a[i][j-1] // east
		}
		if a[i][j] > a[i][j+1] {
			p += a[i][j] - a[i][j+1] // west
		}

		if a[i][j] > a[i-1][j] {
			p += a[i][j] - a[i-1][j] // north
		}

		if a[i][j] > a[i+1][j] {
			p += a[i][j] - a[i+1][j] // south
		}

		sum += p
	}

	ch <- sum
}

func convertRow(buf []byte, row []int, n int) {
	l := len(buf)
	if buf[l-1] == '\n' {
		l-- // remove '\n'
	}

	p := 1
	i := 0
	for i < l && p <= n {
		for i < l && ' ' == buf[i] {
			i++
		}
		d := 0
		for i < l && ' ' != buf[i] {
			d = d*10 + int(buf[i]-'0')
			i++
		}
		row[p] = d
		p++
	}
}
