# https://leetcode.com/problems/balanced-binary-tree/

from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
def isBalanced(node: Optional[TreeNode]) -> bool:
    if node is None: return True
    
    # balanced if it's a leaf
    if node.left is None and node.right is None:
        node.val=1
        return True
    
    leftOk=isBalanced(node.left)
    if not leftOk:
        return False
    
    rightOk=isBalanced(node.right)
    if not rightOk:
        return False
    
    lh=0 if node.left is None else node.left.val
    rh=0 if node.right is None else node.right.val
    
    if lh>rh: 
        node.val=lh+1
    else: 
        node.val=rh+1
    
    return abs(lh-rh)<2