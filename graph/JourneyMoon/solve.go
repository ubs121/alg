package main

import "fmt"

func main() {
	var n, p int

	fmt.Scanf("%d %d", &n, &p)

	uf := MakeUF(n)

	var x, y int
	for i := 0; i < p; i++ {
		x, y = readPair()
		uf.Union(x, y)
	}

	// count connections
	cm := uf.components()
	ans := int64(0)
	for _, m := range cm {
		ans += int64(m * (n - m))
	}
	ans = ans / 2

	fmt.Printf("%d\n", ans)
}

// UF is a Union Find structure
type UF struct {
	id []int // parent reference tree
	sz []int // number of members
}

// MakeUF makes a new Union Find structure
func MakeUF(dim int) *UF {
	uf := new(UF)
	uf.id = make([]int, dim)
	uf.sz = make([]int, dim)

	for i := 0; i < dim; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return uf
}

// Root finds a root of p
func (u *UF) Root(p int) int {
	for p != u.id[p] {
		u.id[p] = u.id[u.id[p]]
		p = u.id[p]
	}
	return p
}

// Connected checks if p and q are connected.
func (u *UF) Connected(p, q int) bool {
	return u.Root(p) == u.Root(q)
}

// Union connects p & q.
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

func (u *UF) components() map[int]int {
	cmap := make(map[int]int) // component id :-> component size
	r := 0
	for i := 0; i < len(u.id); i++ {
		r = u.Root(i)
		cmap[r] = u.sz[r]
	}
	return cmap
}

func readPair() (int, int) {
	var x, y int
	_, err := fmt.Scanf("%d %d\n", &x, &y)

	if err != nil {
		panic(err)
	}

	return x, y
}
