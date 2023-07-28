// https://www.hackerrank.com/challenges/insertion-sort?h_r=internal-search
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	T, _ := readInt(reader)

	var n int
	for t := 0; t < T; t++ {
		n, _ = readInt(reader)

		maxVal := n
		// read input
		a := make([]int, n)
		buf, _ := reader.ReadBytes('\n')

		l := len(buf)
		if buf[l-1] == '\n' {
			l-- // remove '\n'
		}

		p := 0 // index for 'a'
		i := 0
		for i < l && p < n {
			for i < l && ' ' == buf[i] {
				i++
			}
			d := 0
			for i < l && ' ' != buf[i] {
				d = d*10 + int(buf[i]-'0')
				i++
			}
			a[p] = d
			p++

			if maxVal < d {
				maxVal = d
			}
		}
		maxVal++

		biTree := make([]int, maxVal+1) // binary indexed tree
		s := 0
		for i := 0; i < n; i++ {
			v := a[i]
			// get number of lower elements before a[i]
			p = sum(biTree, v)
			s += i - p

			//fmt.Printf("%d - %d\n", v, i-p)

			// increase a[i] by 1
			update(biTree, maxVal, v, 1)
		}

		fmt.Printf("%d\n", s)
	}
}

// Returns sum of a[:index], where 'a' is original array
func sum(biTree []int, index int) (sum int) {
	index = index + 1 // index in biTree[]

	// Traverse ancestors of biTree[index]
	for index > 0 {
		sum += biTree[index]      // add to sum
		index -= index & (-index) // move index to parent node
	}
	return
}

// Updates a node in Binary Index Tree (BITree)
func update(biTree []int, n, index, val int) {
	index = index + 1 // index in biTree[] is 1 more than the index in a[]

	// Traverse all ancestors and add 'val'
	for index <= n {
		biTree[index] += val      // Add 'val' to current node of BI Tree
		index += index & (-index) // Update index to that of parent in update View
	}
}

// read an integer from stdin
func readInt(reader *bufio.Reader) (int, error) {
	buf, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}

	l := len(buf) - 1
	n := 0
	for i := 0; i < l; i++ { // exclude '\n'
		n = n*10 + int(buf[i]-'0')
	}
	return n, nil
}
