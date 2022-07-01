package graph

import "math"

// Graph API, undirected
type Graph struct {
	V   int           // number of vertices
	E   int           // number of edges
	Adj map[int][]int // adjacent nodes
}

// Create a graph with v vertex
func NewGraph(v int) *Graph {
	g := &Graph{}
	g.V = v
	g.E = 0
	g.Adj = make(map[int][]int)
	return g
}

// Add an undirected edge v-u
func (g *Graph) AddEdge(v, u int) {
	// NOTE: validate v and u is in range 0..g.V
	g.E++
	g.Adj[v] = append(g.Adj[v], u)
	g.Adj[u] = append(g.Adj[u], v)
}

// Returns the degree of v
func (g *Graph) Degree(v int) int {
	return len(g.Adj[v])
}

// Computes BFS shortest paths from a source ro every other vertex
type BFS struct {
	marked []bool // marked[v] - is there an src-dest path
	edgeTo []int  // edgeTo[v] - previous edge on shortest src-dest path
	distTo []int  // distTo[v] - number of edges from source
}

// Creates a new BFS on a graph
func NewBFS(g *Graph, src int) *BFS {
	bfs := &BFS{}
	bfs.marked = make([]bool, g.V)
	bfs.distTo = make([]int, g.V)
	bfs.edgeTo = make([]int, g.V)
	bfs.run(g, src)
	return bfs
}

// Breath First Search from a single source 's'
func (bfs *BFS) run(g *Graph, s int) {
	var q []int // queue

	// init distances to max
	for v := 0; v < g.V; v++ {
		bfs.distTo[v] = math.MaxInt64
	}

	bfs.marked[s] = true // label 's' as discovered
	bfs.distTo[s] = 0    // distance is 0
	q = append(q, s)     // enqueue 's'

	for len(q) > 0 {
		// dequeue
		u := q[0]
		q = q[1:]

		// for all edges from u to v
		for _, v := range g.Adj[u] {
			if !bfs.marked[v] {
				bfs.edgeTo[v] = u
				bfs.distTo[v] = bfs.distTo[u] + 1
				bfs.marked[v] = true
				q = append(q, v) // enqueue
			}
		}
	}
}

// Is there a path between source and v
func (bfs *BFS) HasPathTo(v int) bool {
	return bfs.marked[v]
}

// Returns a shortest path between source and v
func (bfs *BFS) PathTo(v int) []int {
	if !bfs.marked[v] {
		return nil
	}

	var path []int
	p := v
	for bfs.distTo[p] != 0 {
		path = append(path, p)
		p = bfs.edgeTo[p]
	}
	path = append(path, p) // source itself
	return path
}

// Depth First Search on a undirected graph
type DFS struct {
	marked []bool // marked[v] - is there an src-dest path
	count  int    // number of vertices connected to source
}

// Creates a new DFS from source on a graph
func NewDFS(g *Graph, src int) *DFS {
	dfs := &DFS{}
	dfs.marked = make([]bool, g.V)
	dfs.run(g, src)
	return dfs
}

// depth first search from 'u'
func (dfs *DFS) run(g *Graph, u int) {
	dfs.count++
	dfs.marked[u] = true

	for _, v := range g.Adj[u] {
		if !dfs.marked[v] {
			dfs.run(g, v)
		}
	}
}

// Is there a path between source and v
func (dfs *DFS) HasPathTo(v int) bool {
	return dfs.marked[v]
}

// Number of vertices connected to the source
func (dfs *DFS) Count() int {
	return dfs.count
}

// Edge-weighted graph of vertices
type EdgeWeightedGraph struct {
	V int // number of vertices
	E int // number of edges
}
