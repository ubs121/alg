
# sequence sum example
def S(seq, i=0):
    if i == len(seq): return 0
    return S(seq, i+1) + seq[i]

# T (n) = T (n - 1) + 1
def T(seq, i=0):
    if i == len(seq): return 1
    return T(seq, i+1) + 1

seq = range(1,101)
print(S(seq))
print(T(seq))