package alg

import (
	"testing"
)

// Binary tree node
type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// Pre-order traversal is to visit the root first. Then traverse the left subtree. Finally, traverse the right subtree
func PreorderTraversal(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}

	// values from pre-order traversal
	var ret []int
	visit := func(e *BinaryTreeNode) {
		ret = append(ret, e.Data)
	}

	var stack []*BinaryTreeNode // processing stack
	stack = append(stack, root) // push root

	for len(stack) > 0 {
		// pop
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visit(e)

		// push right first so that left processed first
		if e.Right != nil {
			stack = append(stack, e.Right)
		}

		// push left
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}

	return ret
}

// In-order traversal is to traverse the left subtree first. Then visit the root. Finally, traverse the right subtree.
func InorderTraversal(root *BinaryTreeNode) []int {
	// values from in-order traversal
	var ret []int
	visit := func(e *BinaryTreeNode) {
		ret = append(ret, e.Data)
	}

	var stack []*BinaryTreeNode // processing stack

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1] // pop
			stack = stack[:len(stack)-1]

			visit(root)

			root = root.Right
		}
	}

	return ret
}

func TestInorderTraversal(t *testing.T) {

}

// Post-order traversal is to traverse the left subtree first. Then traverse the right subtree. Finally, visit the root.
func PostorderTraversal(root *BinaryTreeNode) []int {
	// values from post-order traversal
	var ret []int
	visit := func(e *BinaryTreeNode) {
		ret = append(ret, e.Data)
	}

	var stack []*BinaryTreeNode  // processing stack
	var lastNode *BinaryTreeNode // last node pointer

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			topNode := stack[len(stack)-1] // peek

			if topNode.Right != nil && lastNode != topNode.Right {
				root = topNode.Right
			} else {
				visit(topNode)

				// pop
				lastNode = topNode
				stack = stack[:len(stack)-1]
			}
		}
	}

	return ret
}

// Level-order traversal (breadth first) is to traverse the tree level by level
func LevelorderTraversal(root *BinaryTreeNode) [][]int {

	if root == nil {
		return nil
	}

	// values from level-order traversal, level by level
	var ret [][]int

	var q []*BinaryTreeNode // processing queue
	q = append(q, root)

	for len(q) > 0 {

		// visit all items in the level
		var nextQ []*BinaryTreeNode
		var vals []int
		for _, e := range q {
			vals = append(vals, e.Data)

			if e.Left != nil {
				nextQ = append(nextQ, e.Left)
			}

			if e.Right != nil {
				nextQ = append(nextQ, e.Right)
			}
		}

		ret = append(ret, vals) // collect it
		q = nextQ               // next level queue
	}

	return ret
}

// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
func MaximumDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaximumDepth(root.Left)
	rightDepth := MaximumDepth(root.Right) // can be called concurrently if it's a large tree

	max := leftDepth
	if max < rightDepth {
		max = rightDepth
	}
	return max + 1
}

func hasPathSum(root *BinaryTreeNode, targetSum int) bool {

	return pathSum(root, 0, targetSum)
}

func pathSum(root *BinaryTreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Data+sum == targetSum {
			return true
		}
	}

	return pathSum(root.Left, root.Data+sum, targetSum) ||
		pathSum(root.Right, root.Data+sum, targetSum)
}

func TestHasPathSum(t *testing.T) {
	root := &BinaryTreeNode{Data: 1}
	l1 := &BinaryTreeNode{Data: -2}
	r1 := &BinaryTreeNode{Data: -3}
	root.Left = l1
	root.Right = r1

	l2 := &BinaryTreeNode{Data: 1}
	r2 := &BinaryTreeNode{Data: 3}
	l1.Left = l2
	l1.Right = r2

	l2.Left = &BinaryTreeNode{Data: -1}

	r1.Left = &BinaryTreeNode{Data: -2}

	ans := hasPathSum(root, 3)
	if ans != false {
		t.Errorf("exp: false, got : true")
	}
}

// BinarySearchTreeNode is same as BinaryTreeNode, but behaviour is different
type BinarySearchTreeNode BinaryTreeNode

//TODO: implement insert(), findMin(), findMax()

// AVL (Adelson-Velskii and Landis) binary search tree
type AVLTreeNode struct {
	Left   *AVLTreeNode
	Data   int
	Right  *AVLTreeNode
	Height int
}
