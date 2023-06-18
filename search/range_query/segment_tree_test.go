package search

import (
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)
	fmt.Println(st.getSum(arr, 1, 3))
	st.updateValue(arr, 1, 10)
	fmt.Println(st.getSum(arr, 1, 3))
}
