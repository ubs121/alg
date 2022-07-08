from bisect import bisect

def lis(arr: list[int|float]):
    '''Longest increasing subsequence'''
    end = [] # end values for all lengths
    for val in arr:
        idx = bisect(end, val) # binary search
        if idx == len(end):
            end.append(val)
        else:
            end[idx] = val
    return len(end)

# testing
print(lis([0,1.0,5,2,6,3.0,4]))