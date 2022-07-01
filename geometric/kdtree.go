package geometric

import "sort"

type Node struct {
	x, y  int // location
	left  *Node
	right *Node
}

func (node *Node) isLeaf() bool {
	return node.left == nil && node.right == nil
}

func kdTree(offices [][]int, depth int) *Node {
	if len(offices) == 0 {
		return nil
	}
	axis := depth % 2 // select axis based on depth

	// sort by axis and choose median
	sort.Slice(offices, func(i, j int) bool { return offices[i][axis] < offices[j][axis] })
	median := len(offices) / 2

	// create node and subtrees
	return &Node{
		x:     offices[median][0],
		y:     offices[median][1],
		left:  kdTree(offices[:median], depth+1),
		right: kdTree(offices[median+1:], depth+1),
	}
}
