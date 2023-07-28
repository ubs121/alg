# https://leetcode.com/problems/max-points-on-a-line/
from collections import defaultdict
from typing import List
from cmath import inf

def slope(p:List[int], q: List[int])->float:
    if p[0]==q[0]:
        return -inf if p[1]==q[1] else inf
    else:
        return +0.0 if p[1]==q[1] else (q[1]-p[1])/(q[0]-p[0])

# Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, 
# return the maximum number of points that lie on the same straight line.
class Solution:
    def maxPoints(self, points: List[List[int]]) -> int:
        if len(points)<2:
            return 1 # can't make a line, but there's 1 point

        maxp=2
        slope_cnt=defaultdict(int)
        
        for i, p in enumerate(points):
            if maxp<len(points)-i and i+1<len(points):
                slope_cnt.clear() # reset counter
                # for each point determine slopes they make with 'p'
                # group points by slopes and find max with same slopes
                for q in points[i+1:]:
                    sl=round(slope(p, q), 6) # calculate slope
                    slope_cnt[sl]+=1

                m = max(slope_cnt.values())+1 # +1 is to count 'p' itself
                if maxp < m: # accumulate max
                    maxp=m
        return maxp

points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
#points=[[1,1],[2,2],[3,3]]
sol=Solution()
print(sol.maxPoints(points))
