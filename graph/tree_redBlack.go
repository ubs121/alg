package graph

import "strings"

type Color int

const (
	RED Color = iota
	BLACK
)

// RedBlackTree node
type RbtNode struct {
	Key         string   // key
	Data        string   // data
	Left, Right *RbtNode // left, right subtrees
	Color       Color    // color of parent link
	Size        int      // subtree count
}

type RedBlackBST struct {
	root *RbtNode
}

// Inserts a key-value pair into the tree
func (rbt *RedBlackBST) Put(key string, data string) {
	rbt.root = rbt.put(rbt.root, key, data)
	rbt.root.Color = BLACK // enforces the root node to be black
}

func (rbt *RedBlackBST) put(n *RbtNode, key string, data string) *RbtNode {
	if n == nil {
		return &RbtNode{Key: key, Data: data, Color: RED, Size: 1}
	}

	cmp := strings.Compare(key, n.Key)
	if cmp < 0 {
		n.Left = rbt.put(n.Left, key, data)
	} else if cmp > 0 {
		n.Right = rbt.put(n.Right, key, data)
	} else {
		n.Data = data
	}

	if rbt.IsRed(n.Right) && !rbt.IsRed(n.Left) {
		n = rbt.RotateLeft(n)
	}
	if rbt.IsRed(n.Left) && rbt.IsRed(n.Left.Left) {
		n = rbt.RotateRight(n)
	}
	if rbt.IsRed(n.Left) && rbt.IsRed(n.Right) {
		rbt.FlipColors(n)
	}

	n.Size = rbt.Size(n.Left) + rbt.Size(n.Right) + 1
	return n
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
	return rbt.Size(n.Left) + rbt.Size(n.Right) + 1
}

// Get finds a node with a given key
func (rbt *RedBlackBST) Get(key string) *RbtNode {
	x := rbt.root
	for x != nil {
		cmp := strings.Compare(key, x.Key)
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
func (rbt *RedBlackBST) RotateLeft(n *RbtNode) *RbtNode {
	x := n.Right
	n.Right = x.Left
	x.Left = n
	x.Color = n.Color
	n.Color = RED
	x.Size = n.Size
	n.Size = 1 + rbt.Size(n.Left) + rbt.Size(n.Right)
	return x
}

// RotateRight orients a left-leaning red link to lean right
func (rbt *RedBlackBST) RotateRight(n *RbtNode) *RbtNode {
	x := n.Left
	n.Left = x.Right
	x.Right = n
	x.Color = n.Color
	n.Color = RED
	n.Size = x.Size
	x.Size = 1 + rbt.Size(x.Left) + rbt.Size(x.Right)
	return x
}

// FlipColors recolor to split 4-node (2 red links)
func (rbt *RedBlackBST) FlipColors(n *RbtNode) {
	n.Color = RED
	n.Left.Color = BLACK
	n.Right.Color = BLACK
}
