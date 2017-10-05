package alg

// Union Find structure
type UF struct {
	id []int // parent reference tree
	sz []int // number of members
}

// make set
func MakeSet(dim int) *UF {
	uf := new(UF)
	uf.id = make([]int, dim)
	uf.sz = make([]int, dim)

	for i := 0; i < dim; i++ {
		uf.id[i] = i
	}
	return uf
}

// find root of p
func (u *UF) Root(p int) int {
	for p != u.id[p] {
		u.id[p] = u.id[u.id[p]] // Make every other node in path point to its grandparent (thereby halving path length)
		p = u.id[p]
	}
	return p
}

// check if p and q connected
func (u *UF) Connected(p, q int) bool {
	return u.Root(p) == u.Root(q)
}

// connect p & q
func (u *UF) Union(p, q int) {
	i := u.Root(p)
	j := u.Root(q)

	if i == j {
		return
	}

	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
}
