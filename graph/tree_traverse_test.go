package graph

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeLvl struct {
	TreeNode
	Level int
}

// https://leetcode.com/problems/even-odd-tree/
// A binary tree is named Even-Odd if it meets the following conditions:
//
// The root of the binary tree is at level index 0, its children are at level index 1, their children are at level index 2, etc.
// For every even-indexed level, all nodes at the level have odd integer values in strictly increasing order (from left to right).
// For every odd-indexed level, all nodes at the level have even integer values in strictly decreasing order (from left to right).
// Given the root of a binary tree, return true if the binary tree is Even-Odd, otherwise return false.
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q []TreeNodeLvl
	q = append(q, TreeNodeLvl{*root, 0}) // push the root

	lastElemAtLvl := map[int]int{} // level => last value, can be an array

	for len(q) > 0 {
		node := q[0] // top element

		// validate the even/odd requirement
		if node.Level%2 == node.Val%2 {
			return false // failed
		}

		// validate the order requirement
		if lastVal, exists := lastElemAtLvl[node.Level]; exists {
			if node.Level%2 == 0 {
				if lastVal >= node.Val {
					return false // failed
				}
			} else {
				if lastVal <= node.Val {
					return false // failed
				}
			}
		}

		if node.Left != nil {
			q = append(q, TreeNodeLvl{*node.Left, node.Level + 1})
		}
		if node.Right != nil {
			q = append(q, TreeNodeLvl{*node.Right, node.Level + 1})
		}

		lastElemAtLvl[node.Level] = node.Val
		q = q[1:] // remove 'node'

	}
	return true
}

func TestIfValidEvenOdd(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		exp  bool
	}{
		{
			&TreeNode{1,
				&TreeNode{10,
					&TreeNode{3,
						&TreeNode{12, nil, nil},
						&TreeNode{8, nil, nil},
					},
					nil,
				},
				&TreeNode{4,
					&TreeNode{7,
						&TreeNode{6, nil, nil},
						nil,
					},
					&TreeNode{9,
						nil,
						&TreeNode{2, nil, nil},
					},
				},
			},
			true,
		},
		{
			&TreeNode{5,
				&TreeNode{4,
					&TreeNode{3, nil, nil},
					&TreeNode{3, nil, nil}},
				&TreeNode{2,
					&TreeNode{7, nil, nil},
					nil,
				},
			},
			false,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isEvenOddTree(tc.root)
			if got != tc.exp {
				t.Errorf("exp %v, got %v", tc.exp, got)
			}
		})
	}

}

type msg struct {
	elem int
	lvl  int
}

// Flatten tree by levels
func traverseByLevels(root []any) {
	out := make(chan *msg)

	go func() {
		defer close(out)
		_traverse(root, 0, out)
	}()

	lvls := map[int][]int{}

	for {
		m := <-out
		if m == nil {
			break
		}

		lvls[m.lvl] = append(lvls[m.lvl], m.elem)
	}

	fmt.Printf("Answer %v", lvls)
}

func _traverse(root []any, lvl int, out chan *msg) {

	for _, e := range root {
		switch v := e.(type) {
		case int:
			out <- &msg{v, lvl}
		case []any:
			_traverse(v, lvl+1, out)
		default:
			// interface, so go deeper???
		}
	}

}

func TestTraverseLevels(t *testing.T) {
	root := []any{5, 12, []any{1, []any{8, 10, 2}, 1, 100}, []any{15}, 7}
	traverseByLevels(root)

}
