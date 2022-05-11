from collections import defaultdict

def counting_sort(A: list, key=lambda x: x):
    B, C = [], defaultdict(list)

    for x in A:
        C[key(x)].append(x)

    for k in range(min(C), max(C)+1):
        B.extend(C[k])
    return B

# test
print(counting_sort([2, 1, 3]))
