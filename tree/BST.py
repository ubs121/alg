
import sys
from typing import Optional

# Binary Search Tree
class BstNode:
    left = None
    right = None
    lvl = 1
    def __init__(self, key, val) -> None:
        self.key = key
        self.val = val
    def __str__(self, level=0) -> str:
        ret = "\t"*level+ self.key+"\n"
        if self.left is not None:
            ret += self.left.__str__(level+1)
        if self.right is not None:
            ret += self.right.__str__(level+1)
        return ret

# check if a valid BST
def isValidBST(root: Optional[BstNode]) -> bool:
    maxSize = sys.maxsize
    minSize = -sys.maxsize - 1
    return isBST(root, minSize, maxSize)

def isBST(root: Optional[BstNode], minVal, maxVal:int) -> bool:
    if root is None: return True
    return minVal<root.val and root.val<maxVal and isBST(root.left, minVal, root.val) and isBST(root.right, root.val, maxVal)


# AA-Tree balancing: right rotation
def skew(node):
    if None in [node, node.left]: return node
    if node.left.lvl != node.lvl: return node  # no need skew

    lft = node.left
    node.left = lft.right
    lft.right = node
    return lft

# AA-Tree balancing: Left rotation & level increase
def split(node):
    if None in [node, node.right, node.right.right]: return node
    if node.right.right.lvl != node.lvl: return node

    rgt = node.right
    node.right = rgt.left
    rgt.left = node
    rgt.lvl += 1
    return rgt

# AA-Tree balancing: insert a node
def insert(node, key, val):
    if node is None: return BstNode(key, val)

    if node.key == key: node.val = val
    elif key < node.key:
        node.left = insert(node.left, key, val)
    else:
        node.right = insert(node.right, key, val)

    node = skew(node)  # in case it's backward
    node = split(node)  # in case it's overfull
    return node

tree=None

tree=insert(tree, "d","")
tree=insert(tree, "e","")
tree=insert(tree, "f","")
tree=insert(tree, "a","")
tree=insert(tree, "b","")
tree=insert(tree, "c","")

print(tree)