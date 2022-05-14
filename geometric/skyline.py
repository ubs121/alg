# https://leetcode.com/problems/the-skyline-problem/

# divide and conquer with merge sort
# buildings - list of [x,height,y]
def getSkyline(buildings: list[list[int]]) -> list[list[int]]:
    if len(buildings)==1: return [[buildings[0][0],buildings[0][2]], [buildings[0][1],0]]
        
    mid=int(len(buildings) / 2)
    leftSkyline=getSkyline(buildings[:mid])
    rightSkyline=getSkyline(buildings[mid:])
    return mergeSkylines(leftSkyline, rightSkyline)

def mergeSkylines(left:list[list[int]], right: list[list[int]]) -> list[list[int]]:
    result=[]
    i,j,currHeight1,currHeight2,maxHeight,maxCurr=0,0,0,0,0,0

    while i<len(left) and j<len(right):
        if left[i][0] < right[j][0]:  # minimum of x coordinates
            currHeight1=left[i][1]
            maxCurr=max(currHeight1, currHeight2)
            if maxHeight!=maxCurr: result.append([left[i][0], maxCurr])
            i+=1
        elif left[i][0] > right[j][0]: # minimum of x coordinates
            currHeight2=right[j][1]
            maxCurr=max(currHeight2,currHeight1)
            if maxHeight!=maxCurr: result.append([right[j][0], maxCurr])
            j+=1
        else: # same x position
            currHeight1=left[i][1]
            currHeight2=right[j][1]
            maxCurr=max(currHeight1, currHeight2)
            if maxHeight!=maxCurr:
                if left[i][1]>=right[j][1]: result.append([left[i][0], left[i][1]])
                else: result.append([right[j][0], right[j][1]])
            i+=1
            j+=1
        maxHeight=maxCurr
    
    while i<len(left):
        result.append(left[i])
        i+=1
    while j<len(right):
        result.append(right[j])
        j+=1
    return result

buildings = [[0,2,3],[2,5,3]]
res=getSkyline(buildings)
print(res)