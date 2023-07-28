# list.sort() or Tim sort

from select import partition

def quicksort(arr):
    if len(arr) <= 1: return arr 
    lo, pi, hi = partition(arr)                 # pi is in its place
    return quicksort(lo) + [pi] + quicksort(hi) # Sort lo and hi separately

# testing
print(quicksort([3,2,3,1,2,5]))