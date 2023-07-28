package common

import (
	"fmt"
	"reflect"
	"testing"
)

// Maximum number of consecutive elements preceding arr[i] such that a[j]<a[i], j<i
func FindSpans(arr []int) []int {
	stack := new(IntStack)
	p := -1 // index of the closest greater element

	span := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		// pop all lesser elements
		for !stack.IsEmpty() && arr[i] > arr[stack.Peek()] {
			stack.Pop()
		}

		if stack.IsEmpty() {
			p = -1
		} else {
			p = stack.Peek()
		}
		span[i] = i - p
		stack.Push(i)
	}
	return span
}

func TestFindSpans(t *testing.T) {
	testCases := map[string][]int{
		"6,3,4,5,2": {1, 1, 2, 3, 1},
	}

	for tc, exp := range testCases {
		arr := ParseIntArray(tc)
		got := FindSpans(arr)
		if !reflect.DeepEqual(exp, got) {
			t.Errorf("tc [%s]: exp %v, got %v", tc, exp, got)
		}
	}
}

// Given a string that contains brackets (), [] and {}
// validate if brackets are valid and balanced
func isValidBrackets(str string) bool {
	if len(str) == 0 {
		return true
	}

	var stk []byte

	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '[' || str[i] == '{' {
			stk = append(stk, str[i])
			continue
		}

		if len(stk) == 0 {
			return false
		}
		c := stk[len(stk)-1] // peek. top element

		if str[i] == ')' && c != '(' {
			return false
		} else if str[i] == ']' && c != '[' {
			return false
		} else if str[i] == '}' && c != '{' {
			return false
		}

		stk = stk[:len(stk)-1] // pop, because of correct combo
	}
	return len(stk) == 0
}

func TestIsValidBrackets(t *testing.T) {
	testCases := []struct {
		str string
		exp bool
	}{
		{
			str: "()[]{}",
			exp: true,
		},
		{
			str: "([])",
			exp: true,
		},
		{
			str: "([)]{}",
			exp: false,
		},
		{
			str: "([][}])",
			exp: false,
		},
		{
			str: "()[]{}",
			exp: true,
		},
		{
			str: ")))",
			exp: false,
		},
		{
			str: "{{{",
			exp: false,
		},
	}
	for i, tc := range testCases {
		got := isValidBrackets(tc.str)
		if tc.exp != got {
			t.Errorf("tc %d: exp %v, got %v", i, tc.exp, got)
		}
	}
}

// integer stack
type IntStack []int

// Is empty ?
func (s IntStack) IsEmpty() bool {
	return len(s) == 0
}

// стекийн оройд элемент нэмэх
func (s *IntStack) Push(value int) {
	*s = append(*s, value)
}

// стекийн оройгоос элемент авах
func (s *IntStack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	top := len(*s) - 1
	value := (*s)[top]
	*s = (*s)[:top]
	return value
}

func (s IntStack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s[len(s)-1]
}

func TestStack(t *testing.T) {
	stack := new(IntStack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(1)

	for !stack.IsEmpty() {
		fmt.Printf("%d ", stack.Pop())
	}
}

// https://leetcode.com/problems/min-stack/
type MinStack struct {
	elems [][]int // (elem, min) pairs
	n     int     // number of elements
}

/** initialize your data structure here. */
func NewMinStack() *MinStack {
	return &MinStack{}
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
