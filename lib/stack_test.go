package alg

import (
	"fmt"
	"testing"
)

type IntStack struct {
	elems []int
}

// стекийн урт
func (s *IntStack) Len() int {
	return len(s.elems)
}

// стекийн оройд элемент нэмэх
func (s *IntStack) Push(value int) {
	s.elems = append(s.elems, value)
}

// стекийн оройгоос элемент авах
func (s *IntStack) Pop() int {
	if len(s.elems) == 0 {
		panic("stack is empty")
	}

	top := len(s.elems) - 1
	value := s.elems[top]
	s.elems = s.elems[:top]
	return value
}

func TestStack(t *testing.T) {
	stack := new(IntStack)

	stack.Push(1)
	stack.Push(2)
	stack.Push(1)

	for stack.Len() > 0 {
		fmt.Printf("%d ", stack.Pop())
	}
}
