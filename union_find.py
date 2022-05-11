# find components
def union_find(pairs: list):
    pairs = map(set, pairs)
    unions = []
    for pair in pairs:
        temp = []
        for s in unions:
            if s.isdisjoint(pair):
                temp.append(s) # append group
            else:
                pair = s.union(pair) # join
                
        temp.append(pair)
        unions = temp
    return unions

# testing
if __name__ == '__main__':
    l = [[1, 2], [2, 3], [4, 5], [6, 7], [1, 7]]
    print(union_find(l))