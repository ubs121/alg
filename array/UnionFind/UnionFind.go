package graph

// UF is a Union Find structure (dynamic connectivity)
type UF struct {
	id []int // nodes, id[i] - a parent of 'i'
	sz []int // tree size, number of nodes
}

// MakeUF makes a new Union Find structure
func MakeUF(dim int) *UF {
	uf := new(UF)
	uf.id = make([]int, dim)
	uf.sz = make([]int, dim)

	for i := 0; i < dim; i++ {
		uf.id[i] = i // each node is on its own
		uf.sz[i] = 1 // each node size is 1
	}
	return uf
}

// Root finds a root of 'p'
func (u *UF) Root(p int) int {
	for p != u.id[p] {
		u.id[p] = u.id[u.id[p]] // Make every other node in path point to its grandparent (thereby halving path length)
		p = u.id[p]
	}
	return p
}

// Connected checks if 'p' and 'q' are connected. O(log(n))
func (u *UF) Connected(p, q int) bool {
	return u.Root(p) == u.Root(q)
}

// Union connects p & q. O(log(n))
func (u *UF) Union(p, q int) {
	i := u.Root(p)
	j := u.Root(q)

	if i == j {
		return // already connected
	}

	// weigh sizes and connect to the small one,
	// it makes sure each node is not so far from the root
	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
}

// Components is connected node groups
func (u *UF) Components() map[int]int {
	cmap := make(map[int]int) // component id :-> component size
	r := 0
	for i := 0; i < len(u.id); i++ {
		r = u.Root(i)
		cmap[r] = u.sz[r]
	}
	return cmap
}
