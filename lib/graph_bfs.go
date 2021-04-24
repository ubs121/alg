package alg

type Graph struct {
}

// Graph data
type GraphData []V

// vertex
type V struct {
	v   int   // node id
	Adj []int // adjucent nodes

	// BFS data
	Parent int // parent
	Color  int // White=0, Gray=1, Black=2
	d      int // distance from source/root
}

// Breath First Search
func (g GraphData) BFS(root int) {
	var Q []int
	var u, v int

	Q = append(Q, root)
	for len(Q) > 0 {
		// dequeue
		u = Q[0]
		Q = Q[1:]

		for i := 0; i < len(g[u].Adj); i++ {
			v = g[u].Adj[i]

			if g[v].Color == 0 {
				g[v].Color = 1
				//g[v].d = g[u].d + 1
				//g[v].parent = u

				// enqueue
				Q = append(Q, v)
			} else if g[v].Color == 1 {
				// repeated link
			}
		}

		g[u].Color = 2
	}
}
