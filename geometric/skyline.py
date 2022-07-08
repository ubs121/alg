# https://leetcode.com/problems/the-skyline-problem/
from typing import List

# divide and conquer with merge sort
# buildings - list of [x,height,y]
def get_skyline(buildings: List[List[int]]) -> List[List[int]]:
    if len(buildings)==1: 
        return [[buildings[0][0],buildings[0][2]], [buildings[0][1],0]]
        
    mid=int(len(buildings) / 2)
    left_skyline=get_skyline(buildings[:mid])
    right_skyline=get_skyline(buildings[mid:])
    return merge_skylines(left_skyline, right_skyline)

def merge_skylines(left:List[List[int]], right: List[List[int]]) -> List[List[int]]:
    result=[]
    i,j,curr_height1,curr_height2,max_height,max_curr=0,0,0,0,0,0

    while i<len(left) and j<len(right):
        if left[i][0] < right[j][0]:  # minimum of x coordinates
            curr_height1=left[i][1]
            max_curr=max(curr_height1, curr_height2)
            if max_height!=max_curr: 
                result.append([left[i][0], max_curr])
            i+=1
        elif left[i][0] > right[j][0]: # minimum of x coordinates
            curr_height2=right[j][1]
            max_curr=max(curr_height2,curr_height1)
            if max_height!=max_curr:
                result.append([right[j][0], max_curr])
            j+=1
        else: # same x position
            curr_height1=left[i][1]
            curr_height2=right[j][1]
            max_curr=max(curr_height1, curr_height2)
            if max_height!=max_curr:
                if left[i][1]>=right[j][1]:
                    result.append([left[i][0], left[i][1]])
                else:
                    result.append([right[j][0], right[j][1]])
            i+=1
            j+=1
        max_height=max_curr
    
    while i<len(left):
        result.append(left[i])
        i+=1
    while j<len(right):
        result.append(right[j])
        j+=1
    return result

buildings = [[0,2,3],[2,5,3]]
res=get_skyline(buildings)
print(res)
