# https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
from typing import List

def minOperations(grid: List[List[int]], x: int) -> int:
    nums=list([]) # convert into vector
    for i in range(len(grid)):
        nums.extend(grid[i])
    n=len(nums)
    if n==1: return 0
    
    nums.sort() # sort it
    
    # check if there's a solution
    rem=nums[0]%x
    for i in range(1, n):
        if nums[i]%x!=rem:
            return -1 # impossible
    
    # find median and make it target, it makes one 0 operation
    target=nums[int(n/2)]
    return sum([abs(int((target-elem)/x)) for elem in nums])