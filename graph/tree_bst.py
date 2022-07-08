
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
                max_node=findMax(root.left)
                root.key=max_node.key
                root.val=max_node.val
                root.left=delete(root.left, key)
            else:
                if not root.left: root=root.right
                if not root.right: root=root.left
    return root
