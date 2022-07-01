package numbers

import (
	"fmt"
	"math"
	"testing"
)

func isPrime(n int) bool {
	for j := 3; j < int(math.Sqrt(float64(n))); j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func TestPrime(t *testing.T) {
	n := 1000
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
