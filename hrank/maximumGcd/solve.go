// https://www.hackerrank.com/contests/w34/challenges/maximum-gcd-and-sum
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	a := make([]int, n)
	b := make([]int, n)

	reader := bufio.NewReader(os.Stdin)
	readArray(reader, a)
	readArray(reader, b)

	// sort
	sort.Ints(a)
	sort.Ints(b)

	// TODO: optimize
	sum := 0
	max := 0
	for i := n - 1; i >= 0 && a[i] > max; i-- {
		for j := n - 1; j >= 0 && b[j] > max; j-- {
			m := gcd(a[i], b[j])
			if max < m {
				max = m
				sum = a[i] + b[j]
			}
		}
	}

	fmt.Printf("%d\n", sum)
}

// greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t

		// a, b = b, a%b

		// (a, b)
	}
	return a
}

func readArray(reader *bufio.Reader, arr []int) {
	buf, _ := reader.ReadBytes('\n')

	l := len(buf)
	if buf[l-1] == '\n' {
		l-- // remove '\n'
	}

	j := 0 // index for 'arr'
	i := 0 // index for 'buf'
	for i < l && j < len(arr) {
		for i < l && ' ' == buf[i] {
			i++
		}
		d := 0
		for i < l && ' ' != buf[i] {
			d = d*10 + int(buf[i]-'0')
			i++
		}
		arr[j] = d
		j++
	}
}
