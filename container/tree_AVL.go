package container

// AVL (Adelson-Velskii and Landis) binary search tree
type AVLTreeNode struct {
	Left   *AVLTreeNode
	Data   int
	Right  *AVLTreeNode
	Height int
}
