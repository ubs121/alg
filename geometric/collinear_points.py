# https://coursera.cs.princeton.edu/algs4/assignments/collinear/specification.php
from typing import List
from cmath import inf
from collections import defaultdict

def slope(p:List[int], q: List[int])->float:
    if p[0]==q[0]:
        return -inf if p[1]==q[1] else inf
    else:
        return +0.0 if p[1]==q[1] else (q[1]-p[1])/(q[0]-p[0])
    
# The problem. Given a set of n distinct points in the plane, find every (maximal) line segment that connects a subset of 4 or more of the points.
def collinear_points(points: List[List[int]])->List:
    segments=[]
    p_slopes=defaultdict(list) # slope => [q1,q2,...]
    for i, p in enumerate(points):
        p_slopes.clear() # reset for new 'p'

        # for each point determine the slope it makes with p 
        for q in points[i+1:]:
            sl=round(slope(p, q), 6)
            p_slopes[sl].append(q)

        # check for points with same slope
        for ps in p_slopes.values():
            if len(ps)>2: # found at least 3 (p=1 + q=2) points on same line
                ps.append(p) # add 'p' itself
                segments.append(ps)

    return segments

# test cases
points=[[10000, 0], [0, 10000], [3000, 7000], [7000, 3000], [20000, 21000], [3000, 4000], [14000, 15000], [6000, 7000]]
res=collinear_points(points)
assert(len(res)==2)
print(res)
