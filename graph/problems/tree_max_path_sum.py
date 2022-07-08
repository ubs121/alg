# https://leetcode.com/problems/binary-tree-maximum-path-sum/
from typing import Optional
import math

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        (maxSum, _)=self.calcSums(root)
        return maxSum
    
    
    def calcSums(self, root: Optional[TreeNode]) -> tuple[int,int]:
        if root is None:
            return -math.inf, 0

        (left_max_sum, left_sum)=self.calcSums(root.left)
        (right_max_sum, right_sum)=self.calcSums(root.right)

        left_sum=max(left_sum, 0) # don't include left branch if it's negative
        right_sum=max(right_sum, 0) # don't include the right branch if it's negative

        # (max value, left or right branch sum - can't be both)
        return max(root.val, left_max_sum, right_max_sum, root.val+left_sum, root.val+right_sum, root.val+right_sum+left_sum), max(root.val+left_sum, root.val+right_sum)


def test_solution():
    test_cases = {
        "1": (TreeNode(1, TreeNode(2), TreeNode(3)), 6),
        "2": (TreeNode(-10, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7))), 42),
        "3": (TreeNode(5, TreeNode(4, TreeNode(11, TreeNode(7), TreeNode(2))), TreeNode(8, TreeNode(13), TreeNode(4, TreeNode(1)))), 48),
        "4": (TreeNode(9, TreeNode(6), TreeNode(-3, TreeNode(-6), TreeNode(2, TreeNode(2, TreeNode(-6, TreeNode(-6)), TreeNode(-6))))), 16)
    }
    sol = Solution()
    for tc, (root, exp) in test_cases.items():
        got = sol.maxPathSum(root)
        if got != exp:
            print(f"Test {tc}: exp {exp}, got {got}")
        else:
            print(f"Test {tc}: passed.")

if __name__ == '__main__':
    test_solution()
