package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 55, 77}
	x := 10 // search value

	// to find and return the smallest index i in [0, n) at which func(i) is true
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		// x is present at arr[i]
		fmt.Printf("%d is at %d", x, i)
	} else {
		// x is not present in arr,
		// but i is the index where it would be inserted.
		fmt.Printf("%d is not found, but can be inserted at %d", x, i)
	}
}
