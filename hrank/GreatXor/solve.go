package main

import "fmt"

var pow2 map[int]int

// calculate 2 powers
func initPow() {
	pow2 = make(map[int]int)

	p := 1
	for i := 0; i < 65; i++ {
		pow2[i] = p
		p = p << 1
	}
}

func xorCount(x int) int {
	i := x

	// number of binary digits
	d := 0
	for i > 0 {
		d++
		i = i / 2
	}

	return pow2[d] - x - 1
}

func main() {
	var q, x int
	fmt.Scanf("%d\n", &q)

	initPow()
	for i := 0; i < q; i++ {
		fmt.Scanf("%d\n", &x)
		fmt.Println(xorCount(x))
	}
}
