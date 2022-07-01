# https://leetcode.com/problems/jump-game-ii/
from heapq import heappush, heappop

inf=float('inf')

def jump_dp(nums: list[int]) -> int:
    if len(nums)==0 or nums[0]==0: return inf
    
    jumps=[inf]*len(nums) # jumps[u] - minimum jumps to come at 'u', by default 'inf'
    jumps[0]=0 # first element
    for u in range(1, len(nums)):
        for v in range(0, u):
            if v+nums[v]>=u: # check if 'u' is reachable from 'v'
                if jumps[v]+1<jumps[u]:
                    jumps[u]= jumps[v]+1 # better solution
    return jumps[-1] # answer

# problem: shortest path using "a star" algorithm
def jump_a_star(nums: list[int]) -> int:
    source, target = 0, len(nums)-1
    jumps, pre, Q = {source:0}, {}, [(0, source)] # (cost, pred, node)
    while Q:
        _, u = heappop(Q)
        if u==target: return jumps[target] # reached

        u_nbrs=list(range(u+1,u+nums[u]+1)) # all reachable cells
        for v in u_nbrs:
            if v>target or (nums[v]==0 and v!=target): continue # invalid neighbor
            d=jumps.get(u,inf)+1
            if d<jumps.get(v, inf): # better solution
                jumps[v]=d
                if v not in pre:
                    heappush(Q, (d+1, v)) # (priority,Node)
                pre[v]=u
    return inf # goal was never reached

# problem: range min
def jump_range_min(nums: list[int]) -> int:
    # TODO: pre-compute logarithms
    lg=[0,1]
    for i in range(2, len(nums)+1):
        lg[i]=lg[i/2]+1

    # TODO: go backward and build a sparse tree
    # TODO: find range (in reachable neighbors) minimum , and set min+1
    
    return -1 # goal was never reached

# testing
testCases={
    "tc1":([2,3,1,1,4], 2),
    "tc2":([2,2,0,1,4], 3),
    "tc3":([5,9,3,2,1,0,2,3,3,1,0,0], 3)
}
for tc,(nums,exp) in testCases.items():
    got=jump_dp(nums)
    if got!=exp:
        print("{}: exp {}, got {}".format(tc, exp, got))
    else:
        print("{}: ok".format(tc))