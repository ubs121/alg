# https://leetcode.com/problems/binary-tree-cameras/

from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
class Solution:
    def minCameraCover(self, root: Optional[TreeNode]) -> int:
        return placeCamera(root, True)
    
def placeCamera(node: Optional[TreeNode], isRoot=False) -> int:
    if node is None:
        return 0

    # leaf
    if node.left is None and node.right is None:
        if isRoot:
            node.val=2 # place camera
            return 1
        
        node.val=0
        return 0

    lc=placeCamera(node.left)
    rc=placeCamera(node.right)

    if (node.left and node.left.val==0) or (node.right and node.right.val==0):
        node.val=2 # place camera here
        return lc+rc+1

    if (node.left and node.left.val==2) or (node.right and node.right.val==2):
        node.val=1 # being watched by one of the children, no camera needed
        return lc+rc

    # (node.left and node.left.val==1) or (node.right and node.right.val==1)
    if isRoot:
        node.val=2 # place camera
        return lc+rc+1
    
    node.val=0
    return lc+rc