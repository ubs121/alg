from collections import deque
from typing import Optional

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# DLR
def preorderTraversal(root: Optional[TreeNode]):
    if root:
        print(root.val)
        preorderTraversal(root.left)
        preorderTraversal(root.right)

# LDR
def inorderTraversal(root: Optional[TreeNode]):
    if root:
        inorderTraversal(root.left)
        print(root.val)
        inorderTraversal(root.right)

# LRD
def postorderTraversal(root: Optional[TreeNode]):
    if root:
        postorderTraversal(root.left)
        postorderTraversal(root.right)
        print(root.val)

def levelorderTraversal(root: Optional[TreeNode]):
    if not root: return
    Q=deque([])
    while Q:
        node=Q.popleft()
        print(node.val)
        if node.left: Q.extend(node.left)
        if node.right: Q.extend(node.right)