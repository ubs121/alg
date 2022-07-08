# https://leetcode.com/problems/jump-game-ii/
from typing import List

inf=float('inf')

def jump_dp(nums: List[int]) -> int:
    if len(nums)==0 or nums[0]==0:
        return inf

    jumps=[inf]*len(nums) # jumps[u] - minimum jumps to come at 'u', by default 'inf'
    jumps[0]=0 # first element
    for u in range(1, len(nums)):
        for v in range(0, u):
            if v+nums[v]>=u: # check if 'u' is reachable from 'v'
                if jumps[v]+1<jumps[u]:
                    jumps[u]= jumps[v]+1 # better solution
    return jumps[-1] # answer

# testing
testCases={
    "tc1":([2,3,1,1,4], 2),
    "tc2":([2,2,0,1,4], 3),
    "tc3":([5,9,3,2,1,0,2,3,3,1,0,0], 3)
}
for tc,(nums,exp) in testCases.items():
    got=jump_dp(nums)
    if got!=exp:
        print(f"{tc}: exp {exp}, got {got}")
    else:
        print(f"{tc}: ok")