// http://www.geeksforgeeks.org/binary-indexed-tree-or-fenwick-tree-2
package main

import (
	"fmt"
)

// Returns sum of a[:index], where 'a' is original array
func sum(biTree []int, index int) int {
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
func update(biTree []int, n, index, val int) {
	index = index + 1 // index in biTree[] is 1 more than the index in a[]

	// Traverse all ancestors and add 'val'
	for index <= n {
		biTree[index] += val      // Add 'val' to current node of BI Tree
		index += index & (-index) // Update index to that of parent in update View
	}
}

func main() {
	freq := make([]int, 10)
	n := len(freq)
	biTree := make([]int, n+1)

	fmt.Printf("Sum of elements in arr[0..5] is %d\n", sum(biTree, 5))

	// freq[3] += 1
	update(biTree, n, 3, 1)
	// freq[2] += 1
	update(biTree, n, 2, 1)

	fmt.Printf("Sum of elements in arr[0..5] after update is %d\n", sum(biTree, 5))
}
