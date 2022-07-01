
def mergesort(seq):
    mid = len(seq)//2
    left, right = seq[:mid], seq[mid:]

    if len(left) > 1: left = mergesort(left)
    if len(right) > 1: right = mergesort(right)

    res = []
    while left and right:
        if left[-1] >=right[-1]:
            res.append(left.pop())
        else:
            res.append(right.pop())
    res.reverse()
    return (left or right) + res

# testing
print(mergesort([3,2,3,1,2,5]))