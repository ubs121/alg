package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

func cmp(a, b []byte, canAssume bool) int {
	lim := len(a)
	if lim > len(b) {
		lim = len(b)
	} else {
		if canAssume && len(a) == len(b) {
			return 0
		}
	}

	r := bytes.Compare(a[:lim], b[:lim])
	if r == 0 {
		return len(b) - len(a) // longest
	}

	return r
}

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for t := 0; t < T; t++ {
		var A, B string
		fmt.Scanf("%s", &A)
		fmt.Scanf("%s", &B)

		a := []byte(A)                     // stack a
		b := []byte(B)                     // stack b
		ans := make([]byte, len(a)+len(b)) // answer, destination string

		p := 0
		c := 0
		canAssume := false

		for len(a) > 0 && len(b) > 0 {
			c = cmp(a, b, canAssume)

			if c < 0 {
				ans[p], a = a[0], a[1:]
			} else if c > 0 {
				ans[p], b = b[0], b[1:]
			} else {
				if rand.Intn(2) == 0 {
					ans[p], a = a[0], a[1:]
				} else {
					ans[p], b = b[0], b[1:]
				}

				// after this point we can assume a == b if len(a) == len(b)
				canAssume = true
			}

			p++
		}

		copy(ans[p:], a)
		copy(ans[p:], b)

		fmt.Printf("%s\n", ans)

	}
}
