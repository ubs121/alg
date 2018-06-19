package main

import (
	"fmt"
	"math/rand"
)

const MAX = 1000000000

func gen(n, h int) {
	fmt.Printf("%d\n", n) // number of students
	fmt.Printf("%d\n", h) // mason's height

	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", rand.Intn(h+100))
	}
	fmt.Println()

	for i := 0; i < n-1; i++ {
		fmt.Printf("%d ", rand.Intn(2*h)-h)
	}
	fmt.Println()
}

func main() {
	gen(1000, 20)
}
