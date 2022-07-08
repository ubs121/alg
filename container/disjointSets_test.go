package container

import (
	"strings"
	"testing"
)

// UF is a Union Find structure (dynamic connectivity)
type UF struct {
	id   []int // nodes, id[i] - a parent of 'i'
	size []int // number of nodes
}

// MakeUF makes a new Union Find structure
// NOTE: dimension has to be n+1 if nodes are numbered from 1..n
func MakeUF(dim int) *UF {
	uf := new(UF)
	uf.id = make([]int, dim)
	uf.size = make([]int, dim)

	for i := 0; i < dim; i++ {
		uf.id[i] = i   // each node is on its own
		uf.size[i] = 1 // each node size is 1
	}
	return uf
}

// Finds a root of 'p'
func (u *UF) Root(p int) int {
	for p != u.id[p] {
		u.id[p] = u.id[u.id[p]] // Make every other node in path point to its grandparent, thereby halving path length
		p = u.id[p]
	}
	return p
}

// Checks if 'p' and 'q' are connected. O(log(n))
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

	// weigh sizes and connect to the small tree,
	// it makes sure each node is not so far from the root
	if u.size[i] < u.size[j] {
		u.id[i] = j
		u.size[j] += u.size[i]
	} else {
		u.id[j] = i
		u.size[i] += u.size[j]
	}
}

// Components returns a connected groups
func (u *UF) Components() map[int]int {
	grps := make(map[int]int) // component id :-> component size
	gid := 0                  // group id

	for i := 0; i < len(u.id); i++ {
		gid = u.Root(i)
		grps[gid] = u.size[gid]
	}
	return grps
}

func TestUnionFind(t *testing.T) {

	n := 10         // number of nodes, nodes are numbered from 0..n-1
	uf := MakeUF(n) // create a new UF

	uf.Union(1, 2) // connect 1 and 2
	uf.Union(2, 3) // connect 2 and 3

	if !uf.Connected(1, 3) {
		t.Errorf("expected 1 and 3 are connected, but didn't")
	}
}

/*
https://introcs.cs.princeton.edu/java/24percolation/
Percolation. We model the system as an n-by-n grid of sites. Each site is either blocked or open; open sites are initially empty.
A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites.
If there is a full site in the bottom row, then we say that the system percolates.
*/
func percolates(grid [][]byte) bool {
	w := len(grid)
	h := len(grid[0]) // expecting non-empty array

	// build connections
	n := h * w          // number of nodes
	uf := MakeUF(n + 2) // create a new UF struct, +2 is for extra 2 nodes
	for i := 0; i < h-1; i++ {
		no := i * w
		for j := 0; j < w-1; j++ {
			if grid[i][j] == '1' {
				if grid[i][j+1] == '1' {
					uf.Union(no+j, no+j+1) // right node
				}
				if grid[i+1][j] == '1' {
					uf.Union(no+j, no+j+w) // below node
				}
			}
		}
	}

	start := n
	// connect first row nodes with 'start'
	for j := 0; j < w; j++ {
		uf.Union(start, j)
	}

	end := n + 1
	// connect last row nodes with 'end'
	for j := 0; j < w; j++ {
		uf.Union(w*(h-1)+j, end)
	}

	return uf.Connected(start, end)
}

func TestPercolation(t *testing.T) {
	// test cases of percolation problem
	testCases := map[string]bool{
		`
00110100
11111001
11001100
00111110
10000011
01011100
01101101
10100010
`: false,
		`
00111000
10011111
11100110
00110111
01110110
01000011
10101111
00000100
`: true,
	}

	for tc, exp := range testCases {
		grid := convertGrid(tc)
		got := percolates(grid)
		if exp != got {
			t.Errorf("tc: %s exp %v, got %v", tc, exp, got)
		}
	}
}

func convertGrid(gridStr string) [][]byte {
	lines := strings.Split(gridStr, "\n")
	grid := make([][]byte, 0)
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue // skip empty lines
		}
		grid = append(grid, []byte(line))
	}
	return grid
}
