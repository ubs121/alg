
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
    '''checks if a valid BST'''
    if root is None: return True
    return minVal<root.val and root.val<maxVal and isBST(root.left, minVal, root.val) and isBST(root.right, root.val, maxVal)

# insert an element in BST
def insert(root: Optional[BstNode], key, val):
    if root:
        if key<root.key: root.left=insert(root.left, key, val)
        else: root.right=insert(root.right, key, val)
    else:
        root=BstNode(key, val)
    return root

# finding maximum element (the right-most node)
def findMax(root: Optional[BstNode])->BstNode:
    if not root: return None
    if not root.right: return root
    return findMax(root.right)

# delete an element from BST
def delete(root:Optional[BstNode], key):
    if root:
        if key<root.key:
            root.left=delete(root.left, key)
        elif key>root.key:
            root.right=delete(root.right, key)
        else:
            if root.left and root.right:
                maxNode=findMax(root.left)
                root.key=maxNode.key
                root.val=maxNode.val
                root.left=delete(root.left, key)
            else:
                if not root.left: root=root.right
                if not root.right: root=root.left

# AA-Tree balancing: right rotation
def aa_skew(node):
    if None in [node, node.left]: return node
    if node.left.lvl != node.lvl: return node  # no need skew

    lft = node.left
    node.left = lft.right
    lft.right = node
    return lft

# AA-Tree balancing: Left rotation & level increase
def aa_split(node):
    if None in [node, node.right, node.right.right]: return node
    if node.right.right.lvl != node.lvl: return node

    rgt = node.right
    node.right = rgt.left
    rgt.left = node
    rgt.lvl += 1
    return rgt

# AA-Tree balancing: insert a node
def aa_insert(node, key, val):
    if node is None: return BstNode(key, val)

    if node.key == key: node.val = val
    elif key < node.key: node.left = aa_insert(node.left, key, val)
    else: node.right = aa_insert(node.right, key, val)

    node = aa_skew(node)  # in case it's backward
    node = aa_split(node)  # in case it's overfull
    return node

tree=None
tree=aa_insert(tree, "d","")
tree=aa_insert(tree, "e","")
tree=aa_insert(tree, "f","")
tree=aa_insert(tree, "a","")
tree=aa_insert(tree, "b","")
tree=aa_insert(tree, "c","")

print(tree)