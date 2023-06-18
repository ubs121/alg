// https://cp-algorithms.com/data_structures/segment_tree.html#sum-queries
package search

import (
	"math"
)

type SegmentTree struct {
	tree []int
	n    int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	h := int(math.Ceil(math.Log2(float64(n))))
	treeSize := 2*int(math.Pow(2, float64(h))) - 1
	tree := make([]int, treeSize)
	st := &SegmentTree{tree, n}
	st.buildTree(arr, 0, 0, n-1)
	return st
}

func (st *SegmentTree) buildTree(arr []int, node int, start int, end int) {
	if start == end {
		st.tree[node] = arr[start]
		return
	}
	mid := (start + end) / 2
	leftChild := 2*node + 1
	rightChild := 2*node + 2
	st.buildTree(arr, leftChild, start, mid)
	st.buildTree(arr, rightChild, mid+1, end)
	st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
}

func (st *SegmentTree) updateValue(arr []int, index int, value int) {
	diff := value - arr[index]
	arr[index] = value
	st.updateValueUtil(0, 0, st.n-1, index, diff)
}

func (st *SegmentTree) updateValueUtil(node int, start int, end int, index int, diff int) {
	if index < start || index > end {
		return
	}
	st.tree[node] = st.tree[node] + diff
	if start != end {
		mid := (start + end) / 2
		leftChild := 2*node + 1
		rightChild := 2*node + 2
		st.updateValueUtil(leftChild, start, mid, index, diff)
		st.updateValueUtil(rightChild, mid+1, end, index, diff)
	}
}

func (st *SegmentTree) getSum(arr []int, queryStart int, queryEnd int) int {
	if queryStart < 0 || queryEnd > st.n-1 || queryStart > queryEnd {
		return -1
	}
	return st.getSumUtil(0, 0, st.n-1, queryStart, queryEnd)
}

func (st *SegmentTree) getSumUtil(node int, start int, end int, queryStart int, queryEnd int) int {
	if queryStart <= start && queryEnd >= end {
		return st.tree[node]
	}
	if end < queryStart || start > queryEnd {
		return 0
	}
	mid := (start + end) / 2
	leftChild := 2*node + 1
	rightChild := 2*node + 2
	return st.getSumUtil(leftChild, start, mid, queryStart, queryEnd) + st.getSumUtil(rightChild, mid+1, end, queryStart, queryEnd)
}
