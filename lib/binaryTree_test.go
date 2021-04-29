package alg

import "testing"

// Binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Pre-order traversal is to visit the root first. Then traverse the left subtree. Finally, traverse the right subtree
func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// values from pre-order traversal
	var ret []int
	visit := func(e *TreeNode) {
		ret = append(ret, e.Val)
	}

	var stack []*TreeNode       // processing stack
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
func InorderTraversal(root *TreeNode) []int {
	// values from in-order traversal
	var ret []int
	visit := func(e *TreeNode) {
		ret = append(ret, e.Val)
	}

	var stack []*TreeNode // processing stack

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
func PostorderTraversal(root *TreeNode) []int {
	// values from post-order traversal
	var ret []int
	visit := func(e *TreeNode) {
		ret = append(ret, e.Val)
	}

	var stack []*TreeNode  // processing stack
	var lastNode *TreeNode // last node pointer

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
func LevelorderTraversal(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	// values from level-order traversal, level by level
	var ret [][]int

	var q []*TreeNode // processing queue
	q = append(q, root)

	for len(q) > 0 {

		// visit all items in the level
		var nextQ []*TreeNode
		var vals []int
		for _, e := range q {
			vals = append(vals, e.Val)

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
func MaximumDepth(root *TreeNode) int {
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

func hasPathSum(root *TreeNode, targetSum int) bool {

	return pathSum(root, 0, targetSum)
}

func pathSum(root *TreeNode, sum, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val+sum == targetSum {
			return true
		}
	}

	return pathSum(root.Left, root.Val+sum, targetSum) ||
		pathSum(root.Right, root.Val+sum, targetSum)
}

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{Val: 1}
	l1 := &TreeNode{Val: -2}
	r1 := &TreeNode{Val: -3}
	root.Left = l1
	root.Right = r1

	l2 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 3}
	l1.Left = l2
	l1.Right = r2

	l2.Left = &TreeNode{Val: -1}

	r1.Left = &TreeNode{Val: -2}

	ans := hasPathSum(root, 3)
	if ans != false {
		t.Errorf("exp: false, got : true")
	}
}
