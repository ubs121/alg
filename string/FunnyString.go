package main

import (
	"fmt"
	"math"
)

func isFunny(s string) bool {
	ss := []int(s)
	n := len(ss)

	for i := 1; i < n/2+1; i++ {
		if math.Abs(float64(ss[i]-ss[i-1])) != math.Abs(float64(ss[n-i]-ss[n-i-1])) {
			return false
		}
	}
	return true
}

func main() {
	var t int
	fmt.Scanf("%d", &t)

	var s string
	for i := 0; i < t; i++ {
		fmt.Scanf("%s", &s)
		if isFunny(s) {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
	}
}
