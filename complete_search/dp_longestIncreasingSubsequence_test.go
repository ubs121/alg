package search

import (
	"alg/numbers"
	"testing"
)

func solveLongestIncSequence(arr []int) {
	n := len(arr)
	length := make([]int, n)

	// TODO: optimize it
	for k := 0; k < n; k++ {
		length[k] = 1
		for i := 0; i < k; i++ {
			if arr[i] < arr[k] {
				length[k] = numbers.Max(length[k], length[i]+1)
			}
		}
	}
}

func TestLongestIncSequence(t *testing.T) {

}
