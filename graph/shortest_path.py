from heapq import heappush, heappop
# https://algs4.cs.princeton.edu/44sp/

inf = float('inf')

# relaxation operation
def relax(W, u, v, dist, prev):
    d = dist.get(u,inf) + W[u][v]
    if d < dist.get(v,inf):
        dist[v], prev[v] = d, u
        return True # there is a change

# Dijkstra’s Algorithm, 
# G - dict of dicts representation
# https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
def dijkstra(G, source):
    dist, prev, Q, S = {source:0}, {}, [(0,source)], set()
    while Q:
        _, u = heappop(Q)
        if u in S: continue
        S.add(u)
        for v in G[u]:
            relax(G, u, v, dist, prev)
            heappush(Q, (dist[v], v))
    return dist, prev

# Topological sort of a Directed, Acyclic Graph
def topsort(G):
    count = dict((u, 0) for u in G)
    for u in G:
        for v in G[u]:
            count[v] += 1

    Q = [u for u in G if count[u] == 0]
    S = []
    while Q:
        u = Q.pop()
        S.append(u)
        for v in G[u]:
            count[v] -= 1
            if count[v] == 0:
                Q.append(v)
    return S

# DAG shortest path
def dag_sp(W, source, target):
    dist = {u:float('inf') for u in W}
    dist[source] = 0
    for u in topsort(W):
        if u == target: break
        for v in W[u]:
            dist[v] = min(dist[v], dist[u] + W[u][v])
    return dist[target]

# A* algorithm
# https://en.wikipedia.org/wiki/A*_search_algorithm
# h is the heuristic function. h(n) estimates the cost to reach goal from node n.
def a_star(G, source, target, h):
    prev, Q = {}, [(h(source), None, source)]
    while Q:
        d, p, u = heappop(Q)
        if u in prev: continue
        prev[u] = p
        if u == target: return d - h(target), prev
        for v in G[u]:
            w = G[u][v] - h(u) + h(v)
            heappush(Q, (d + w, u, v))
    return inf, None # goal was never reached