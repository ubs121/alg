package alg

import "strings"

const RED = 1
const BLACK = 0

// RedBlackTree node
type RbtNode struct {
	Key         string   // key
	Data        string   // data
	Left, Right *RbtNode // left, right subtrees
	Color       int      // color of parent link
	Size        int      // subtree count
}

type RedBlackBST struct {
	root *RbtNode
}

// IsRed checks if n is red; false if n is null ?
func (rbt *RedBlackBST) IsRed(n *RbtNode) bool {
	if n == nil {
		return false
	}
	return n.Color == RED
}

// Size returns a number of nodes in subtree rooted at n; 0 if n is null
func (rbt *RedBlackBST) Size(n *RbtNode) int {
	if n == nil {
		return 0
	}
	return n.Size
}

// Get finds a node with a given key
func (rbt *RedBlackBST) Get(key string) *RbtNode {
	x := rbt.root
	for x != nil {
		cmp := strings.Compare(x.Key, key)
		if cmp < 0 {
			x = x.Left
		} else if cmp > 0 {
			x = x.Right
		} else { // cmp == 0
			return x
		}
	}

	return nil
}

// RotateLeft orients a right-leaning red link to lean left
func RotateLeft(n *RbtNode) *RbtNode {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.Color = n.Color
	n.Color = RED
	return x
}

// RotateRight orients a left-leaning red link to lean right
func RotateRight(n *RbtNode) *RbtNode {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.Color = n.Color
	n.Color = RED
	return x
}

// FlipColors recolor to split 4-node (2 red links)
func FlipColors(n *RbtNode) {
	n.Color = RED
	n.Left.Color = BLACK
	n.Right.Color = BLACK
}
