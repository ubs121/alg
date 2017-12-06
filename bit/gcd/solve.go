// https://www.hackerrank.com/contests/w34/challenges/maximum-gcd-and-sum
package main

import (
	"bufio"
	"fmt"
	"os"
)

type set map[int]bool

// greatest common divisor
func gcd(a, b int) int {
    for b != 0 {
		t := b
		b = a % b
		a = t

		// (a, b)
	}
 	return a
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	a:=make(set) // set A
	b:=make(set) // set B
	
	reader := bufio.NewReader(os.Stdin)
	readSet(reader, a)
	readSet(reader, b)

	for k := range a {
		
		if k%2 == 0 { // k is even
			u:=k
			for u > 0 {
				if b[u] {
					// (k, u) is maximum
					break
				}
				u = u / 2
			}
		} else { // k is odd

		}
	}

	fmt.Printf("%d\n", gcd(5, 2))
}

func readSet(reader *bufio.Reader, m set) {
	buf, _ := reader.ReadBytes('\n')
	
	l := len(buf)
	if buf[l-1] == '\n' {
		l-- // remove '\n'
	}

	i := 0 // index for 'buf'
	for i < l {
		for i < l && ' ' == buf[i] {
			i++
		}
		d := 0
		for i < l && ' ' != buf[i] {
			d = d*10 + int(buf[i]-'0')
			i++
		}
		m[d] = true // d exists
	}
}
