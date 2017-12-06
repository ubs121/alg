package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n, m int // matrix size
	fmt.Scanf("%d %d\n", &n, &m)

	distTo := make([]int, m) // distTo[v] = longest max to 'v'
	lm := make([]int, m)     // lm[i] == left max of 'i'
	lmtop := make([]int, m)  // lmtop[i] == left max with top
	rm := make([]int, m)     // rm[j] == right max of 'j'
	rmtop := make([]int, m)  // rmtop[j] == right max with top

	reader := bufio.NewReader(os.Stdin)
	for c := 0; c < n; c++ {
		buf, _ := reader.ReadBytes('\n')
		arr := convertArray(buf, m) // convert buffer into array

		lm[0] = arr[0]
		lmtop[0] = arr[0] + distTo[0]
		for i := 1; i < m; i++ {
			lm[i] = max(lm[i-1]+arr[i], arr[i]) // kadane's algorithm
			lmtop[i] = max(lmtop[i-1]+arr[i], lm[i]+distTo[i])
		}

		rm[m-1] = arr[m-1]
		rmtop[m-1] = arr[m-1] + distTo[m-1]
		for i := m - 1 - 1; i >= 0; i-- {
			rm[i] = max(rm[i+1]+arr[i], arr[i])
			rmtop[i] = max(rmtop[i+1]+arr[i], rm[i]+distTo[i])
		}

		/*
			let be 'j' is ending cell, 'i' is beginning cell

				distTo[j] == max( lm(i) + sum([i:j]) + rm(j) ), for some 'i'

				lmtop[j] = lm[j] + sum([i:j]), where i < j
				rmtop[j] = rm[j] + sum([j:i]), where i > j

				so, distTo[j] = max(lmtop[j]+rm[j]-arr[j], rmtop[j]+lm[j]-arr[j])
		*/
		for j := 0; j < m; j++ {
			// accumulate longest path
			distTo[j] = max(lmtop[j]+rm[j]-arr[j], rmtop[j]+lm[j]-arr[j])

			//fmt.Printf("(%d<-%d,%d,%d,%d) ", distTo[i], lm[i], lmtop[i], rm[i], rmtop[i])
		}

		//fmt.Println()
	}

	gMax := math.MinInt32
	for v := 0; v < m; v++ {
		if gMax < distTo[v] {
			gMax = distTo[v]
		}
	}
	fmt.Printf("%d\n", gMax)
}

func convertArray(buf []byte, n int) []int {
	l := len(buf)
	if buf[l-1] == '\n' {
		l-- // remove '\n'
	}

	p := 0       // index for 'arr'
	i := 0       // index for 'buf'
	d := 0       // temp number
	neg := false // negative sign

	arr := make([]int, n)

	for i < l && p < n {
		for i < l && ' ' == buf[i] {
			i++
		}

		d = 0
		neg = false
		if buf[i] == '-' {
			neg = true
			i++
		}

		for i < l && ' ' != buf[i] {
			d = d*10 + int(buf[i]-'0')
			i++
		}
		if neg {
			d = -1 * d
		}
		arr[p] = d
		p++
	}

	return arr
}
