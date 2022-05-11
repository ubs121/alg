
from heapq import heappush, heappop

Q=[]
heappush(Q, 15)
heappush(Q, 1)
heappush(Q, 21)
heappush(Q, 3)

print([heappop(Q) for i in range(len(Q))])