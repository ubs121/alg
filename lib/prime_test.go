package alg

import (
	"fmt"
	"math"
	"testing"
)

func TestPrime(t *testing.T) {
	n := 1000
	for i := 3; i < n; i += 2 {
		isPrime := true
		for j := 3; j < int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", i)
		}
	}
}
