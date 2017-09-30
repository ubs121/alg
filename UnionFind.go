// Union Find structure
type UF struct {
	id []int
}

func createUF(dim int) *UF {
	uf := new(UF)
	uf.id = make([]int, dim)
	return uf
}

func (u *UF) root(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *UF) connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *UF) union(p, q int) {
	i := u.root(p)
	j := u.root(q)
	u.id[i] = j
}