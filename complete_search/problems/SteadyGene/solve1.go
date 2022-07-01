// https://www.hackerrank.com/challenges/bear-and-steady-gene

package main

import (
	"bufio"
	"fmt"
	"os"
)

// quadrants, indexed as G=0, A=1, C=2, T=3
type Quad [4]int

// Quad index values
const (
	G = 0
	A = 1
	C = 2
	T = 3
)

func main() {

	file, _ := os.Open("input03.txt")
	defer file.Close()

	buf := make([]byte, 500100)
	scanner := bufio.NewScanner(file) // file, os.Stdin
	scanner.Buffer(buf, 500100)

	var n int // 4<=n<=500'000
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &n)

	scanner.Scan()
	geneStr := scanner.Text() // read gene string

	// required number of letters to be 'steady'
	limit := n / 4

	// count letters, preCount[i] is number of each letters before 'i'
	preCount := make([]Quad, n+1)
	for i := 0; i < n; i++ {
		// count each letter
		switch geneStr[i] {
		case 'G':
			preCount[i][G]++
		case 'A':
			preCount[i][A]++
		case 'C':
			preCount[i][C]++
		case 'T':
			preCount[i][T]++
		}

		preCount[i+1] = preCount[i] // copy to 'i+1' position
	}

	// find the first location that has over limit
	l := 0
	for i := 0; i < n; i++ {
		c := 0
		switch geneStr[i] {
		case 'G':
			c = preCount[i][G]
		case 'A':
			c = preCount[i][A]
		case 'C':
			c = preCount[i][C]
		case 'T':
			c = preCount[i][T]
		}
		if c > limit {
			l = i
			break
		}
	}

	total := preCount[n]
	over := 0 // overflow amount, or minimum length to be replaced
	for p := 0; p < 4; p++ {
		if total[p] > limit {
			over += total[p] - limit
		}
	}

	// find the minimum distance
	min := n
	for i := l - 1; i >= 0; i-- {
		lo := i + over - 1
		hi := n

		if hi > i+min {
			hi = i + min
		}

		// find j between [lo, hi] using binary search
		j := lo
		for lo < hi-1 {
			j = (lo + hi) / 2

			// check if i+j <= limit
			if preCount[i][G]+total[G]-preCount[j][G] <= limit &&
				preCount[i][A]+total[A]-preCount[j][A] <= limit &&
				preCount[i][C]+total[C]-preCount[j][C] <= limit &&
				preCount[i][T]+total[T]-preCount[j][T] <= limit {

				hi = j // j=(lo+j)/2

				if min > j-i {
					min = j - i
				}
			} else {
				lo = j // j=(j+hi)/2
			}
		}
	}

	if over > 0 {
		fmt.Println(min)
	} else {
		fmt.Println(0)
	}

}
