// http://www.geeksforgeeks.org/binary-indexed-tree-or-fenwick-tree-2
package search

import (
	"fmt"
	"testing"
)

// binary indexed tree
type BiTree []int

// Returns sum of a[:index], where 'a' is original array
func (biTree BiTree) Sum(index int) int {
	sum := 0
	index = index + 1 // index in biTree[]

	// Traverse ancestors of biTree[index]
	for index > 0 {
		sum += biTree[index]      // add to sum
		index -= index & (-index) // move index to parent node
	}
	return sum
}

// Updates a node in Binary Index Tree (BITree)
func (biTree BiTree) Update(index, val int) {
	index = index + 1 // index in biTree[] is 1 more than the index in a[]

	// Traverse all ancestors and add 'val'
	for index < len(biTree) {
		biTree[index] += val      // Add 'val' to current node of BI Tree
		index += index & (-index) // Update index to that of parent in update View
	}
}

func TestBIT(t *testing.T) {
	freq := make([]int, 10)
	n := len(freq)
	biTree := make(BiTree, n+1) // n+1

	fmt.Printf("Sum of elements in arr[0..5] is %d\n", biTree.Sum(5))

	// freq[2] += 1
	biTree.Update(2, 1)
	// freq[3] += 1
	biTree.Update(3, 1)

	fmt.Printf("Sum of elements in arr[0..5] after update is %d\n", biTree.Sum(3))
}
