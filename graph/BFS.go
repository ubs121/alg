package graph

// vertex
type V struct {
	v   int   // node id
	adj []int // adjucent nodes

	// BFS data
	parent int // parent
	color  int // White=0, Gray=1, Black=2
	d      int // distance from source/root
}
type Data []V

// BFS
func BFS(data Data, root int) {
	var Q []int
	var u, v int

	Q = append(Q, root)
	for len(Q) > 0 {
		// dequeue
		u = Q[0]
		Q = Q[1:]

		for i := 0; i < len(data[u].adj); i++ {
			v = data[u].adj[i]

			if data[v].color == 0 {
				data[v].color = 1
				//data[v].d = data[u].d + 1
				//data[v].parent = u

				// enqueue
				Q = append(Q, v)
			} else if data[v].color == 1 {
				// repeated link
			}
		}

		data[u].color = 2
	}
}