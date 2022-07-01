package container

import (
	"fmt"
	"reflect"
	"testing"
)

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
