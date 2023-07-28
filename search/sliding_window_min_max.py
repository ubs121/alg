# https://leetcode.com/problems/sliding-window-maximum/
from typing import List

def maxSlidingWindow(nums: List[int], k: int) -> List[int]:
    maxWindow, Q = [], []
    l,r=0,0 # window [l:r+1]
    while r<len(nums):
        # remove elements outside window
        if Q and Q[0][0]<l: Q.pop(0)

        # remove elements smaller than a new right
        while Q and Q[-1][1]<=nums[r]:
            Q.pop()
        
        # add new element to the end of queue       
        Q.append((r,nums[r]))
        
        r+=1
        if r>=k: # beyond this the window is correct
            maxWindow.append(Q[0][1]) # add into window max
            l+=1
    return maxWindow

arr=[-7,-8,7,5,7,1,6,0]
k=4

# # arr=[9,10,9,-7,-4,-8,2,-6]
# k=5

print(maxSlidingWindow(arr, k))