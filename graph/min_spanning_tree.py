# https://algs4.cs.princeton.edu/43mst/

from heapq import heappop, heappush

# Prim's algorithm for finding a minimum spanning tree
# G - an undirected graph
def prim(G, source):
    prev,Q={}, [(0, None, source)] # (weight, node, )
    while Q:
        _,p,u=heappop(Q)
        if u in prev: continue
        prev[u]=p
        for v, w in G[u].items():
            heappush(Q, (w, u, v))
    return prev