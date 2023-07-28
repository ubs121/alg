

def partition(arr):
    pi, arr = arr[0], arr[1:]
    lo = [x for x in arr if x <= pi]
    hi = [x for x in arr if x > pi]
    return lo, pi, hi

# returns the kth smallest element 
def select(arr, k):
    lo, pi, hi = partition(arr)
    m = len(lo)
    if m == k: return pi
    elif m < k: return select(hi, k-m-1)
    else: return select(lo, k)

# test
print(select([2,1,3,4,0], 2))