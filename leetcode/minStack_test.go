package main

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	elems [][]int // (elem, min) pairs
	n     int     // number of elements
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (st *MinStack) Push(val int) {
	if st.n == 0 {
		st.elems = append(st.elems, []int{val, val})
	} else {
		min := st.GetMin()
		if min > val {
			min = val
		}
		st.elems = append(st.elems, []int{val, min})
	}
	st.n++
}

func (st *MinStack) Pop() {
	// always be called on non-empty stacks.
	st.n--
	st.elems = st.elems[:st.n]
}

func (st *MinStack) Top() int {
	// always be called on non-empty stacks.
	return st.elems[st.n-1][0]
}

func (st *MinStack) GetMin() int {
	// always be called on non-empty stacks.
	return st.elems[st.n-1][1]
}
