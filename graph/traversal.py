from collections import deque

# G - a graph represented using adjacency sets
def bfs(G, source):
    '''Iterative Breath First Search (BFS)'''
    parent, Q = {source:None}, deque([source])
    while Q:
        u=Q.popleft() # dequeue
        for v in G[u]:
            if v in parent: continue # already has parent
            parent[v]=u
            Q.append(v)
    return parent


def dfs(G, source):
    '''Iterative Depth-First Search (DFS)'''
    visited, Q = set(), []
    Q.append(source)
    while Q:
        u = Q.pop()
        if u in visited: continue
        visited.add(u)
        Q.extend(G[u])

# general graph traverse
def traverse(G, source):
    visited, Q = set(), set()
    Q.add(source)
    while Q:
        u = Q.pop()
        if u in visited: continue
        visited.add(u)
        for v in G[u]:
            Q.add(v)

if __name__ == '__main__':
    G = {0: [1, 2], 1: [2], 2: [3], 3: [1, 2]}
    print("BFS(G,0): ")
    print(bfs(G, 0))