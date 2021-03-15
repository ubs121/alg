package sort

import (
	"sort"
)

// NOTE: this is a reference implementation of the count sorting,
// it may be useless if each element is unique
func countSort(arr []int) {
	n := len(arr) // number of elements

	// count each elements, could use array instead of map
	countMap := make(map[int]int) // [key->count]
	for i := 0; i < n; i++ {
		countMap[arr[i]]++
	}

	// sort keys
	var keys []int
	for k := range countMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// do prefix sum for each key
	for i := 1; i < len(keys); i++ {
		k := keys[i]
		countMap[k] += countMap[keys[i-1]]
	}

	// place each elements in order
	B := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		B[countMap[arr[i]]-1] = arr[i]
		countMap[arr[i]]--
	}

	// copy back to 'arr'
	copy(arr, B)
}
