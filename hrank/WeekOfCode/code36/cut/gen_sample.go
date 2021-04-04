package main

import (
	"fmt"
	"math/rand"
)

func gen(n, m, k int) {
	fmt.Printf("%d %d %d\n", n, m, k)
	var r int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r = rand.Intn(10000) - 5000
			fmt.Printf("%d ", r)
		}
		fmt.Println()
	}
}

// func main() {
// 	gen(380, 380, 10)
// }
