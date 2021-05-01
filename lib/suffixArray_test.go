package alg

import (
	"fmt"
	"index/suffixarray"
	"testing"
)

func TestSuffixArray(t *testing.T) {
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}
}
