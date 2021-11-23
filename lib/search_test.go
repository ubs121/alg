package alg

import (
	"fmt"
	"index/suffixarray"
	"sort"
	"testing"
)

func TestSuffixArray(t *testing.T) {
	// create an index for data
	index := suffixarray.New([]byte("banana"))

	// lookup
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}
}

// sort.Search()
func TestBinarySearch(t *testing.T) {
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

func TestSearchInts(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	am := len(a) / 2
	bm := len(b) / 2
	amInd := sort.SearchInts(b, a[am])
	bmInd := sort.SearchInts(a, b[bm])

	exp := 2
	if amInd != exp {
		t.Errorf("exp %d, got %d", exp, amInd)
	}

	exp = 1
	if bmInd != exp {
		t.Errorf("exp %d, got %d", exp, bmInd)
	}
}
