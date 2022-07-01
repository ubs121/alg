# https://leetcode.com/problems/rectangle-area/

def computeArea(ax1: int, ay1: int, ax2: int, ay2: int, bx1: int, by1: int, bx2: int, by2: int) -> int:
    total_area=(ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)

    if ax1 <= bx1 < ax2: # (r1,r2)
        if ay2<=by1 or ay1>=by2:  return total_area # no intersection
    elif bx1 <= ax1 < bx2: # (r2,r1)
        if by2<=ay1 or by1>=ay2: return total_area # no intersection
    else:
        return total_area # no intersection

    # find intersection lines
    y1=max(ay1,by1)
    y2=min(ay2,by2)
    x1=max(ax1,bx1)
    x2=min(ax2,bx2)
    inter_area=(y2-y1)*(x2-x1)

    return total_area-inter_area