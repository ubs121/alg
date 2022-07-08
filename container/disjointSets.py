# find components
from typing import List

def union_find(pairs: List):
    pairs = map(set, pairs)
    groups = []
    for pair in pairs:
        temp1 = [] # temp unions
        for grp in groups:
            if grp.isdisjoint(pair):
                temp1.append(grp) # append back
            else:
                pair = grp.union(pair) # join
                
        temp1.append(pair)
        groups = temp1
    return groups

# testing
if __name__ == '__main__':
    l = [[1, 2], [2, 3], [4, 5], [6, 7], [1, 7]]
    print(union_find(l))
    