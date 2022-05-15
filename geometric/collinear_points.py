# https://coursera.cs.princeton.edu/algs4/assignments/collinear/specification.php
# The problem. Given a set of n distinct points in the plane, find every (maximal) line segment that connects a subset of 4 or more of the points.

from cmath import inf

# points in the plane
class Point:
    def __init__(self, x=None, y=None):
        self.x = x
        self.y = y
    def slopeTo(self, that:'Point')->float:
        if that.x==self.x:
            if that.y==self.y: return -inf
            else: return inf
        else:
            if that.y == self.y: return +0.0
            else: return (that.y-self.y)/(that.x-self.x)
    def __repr__(self):
        return "({},{})".format(self.x,self.y)
   
    
# returns a list of collinear segments
def collinear_points(points: list[Point])->list[list[Point]]:
    uniqueLines=dict() # point->slope
    def markAdded(p: Point, slope):
        if p not in uniqueLines: uniqueLines[p]={}
        uniqueLines[p][slope]=True

    def isAdded(p: Point, slope)->bool:
        return (p in uniqueLines and slope in uniqueLines[p])

    result=list()
    for p in points:
        pSlopes=dict()
        # for each other point q, determine the slope it makes with p 
        for q in points: 
            if q==p: continue
            sl=hash(round(p.slopeTo(q), 6))
            if sl not in pSlopes: pSlopes[sl]=[q]
            else: pSlopes[sl].append(q) # slope => [q1,q2,...]

        for sl, seg in pSlopes.items():
            if len(seg)>1: # found at least 3 points that make 2 pairs
                if isAdded(p, sl): continue
                seg.append(p)
                for q in seg: markAdded(q, sl)
                result.append(seg)

    return result

# test cases
points=[Point(10000, 0), Point(0, 10000), Point(3000, 7000), Point(7000, 3000), Point(20000, 21000), Point(3000, 4000), Point(14000, 15000), Point(6000, 7000)]
res=collinear_points(points)
assert(len(res)==2)
print(res)