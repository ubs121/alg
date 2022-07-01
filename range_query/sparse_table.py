# https://cp-algorithms.com/data_structures/sparse-table.html#range-minimum-queries-rmq

from typing import List
import math

def build_sparse_table(nums : List[int]) -> List[List[int]]:
    # required # of columns = log2 (n) + 1
    sz = math.ceil(math.log2(len(nums))) + 1

    # Create a sparse table of size [n][sz]
    st = [0] * len(nums)
    for i in range(len(nums)) :
        st[i] = [0] * sz
    
    rows = len(nums)
    cols = len(st[0])

    # Fill base case values
    for i in range(rows) :
        st[i][0] = nums[i]

    for c in range(1, cols+1) :
        _range = (1<<c)
        r = 0
        while (r + _range <= rows) :
            # Values in the current column are derived from the values in the previous column.
            st[r][c] = min(st[r][c-1], st[r+(1<<(c-1))][c-1])
            r += 1
    
    return st

def rangeMin(left : int, right : int, sparse_table : List[List[int]]) -> int:
    # Find the biggest block of size 2^p that fits in the range [left:right].
    power_of_2 = int (math.log2(right + 1 - left))
    return min(sparse_table[left][power_of_2], sparse_table[right + 1 - ( 1 << power_of_2)][power_of_2])

# testing
if __name__ == "__main__" :
    arr=[4,6,8,7,3,2,9,5,1]
    sparse_table=build_sparse_table(arr)
    l,r=0,4
    print(arr)
    print("range_min({}:{})=".format(l,r), rangeMin(l,r, sparse_table))