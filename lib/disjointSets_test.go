package alg

import "testing"

func TestUnionFind(t *testing.T) {

	n := 10         // number of nodes, nodes are numbered from 0..n-1
	uf := MakeUF(n) // create a new UF

	uf.Union(1, 2) // connect 1 and 2
	uf.Union(2, 3) // connect 2 and 3

	if !uf.Connected(1, 3) {
		t.Errorf("expected 1 and 3 are connected, but didn't")
	}
}

func TestPercolation(t *testing.T) {
	// TODO: test cases of percolation problem
}

/*
https://introcs.cs.princeton.edu/java/24percolation/
Percolation. We model the system as an n-by-n grid of sites. Each site is either blocked or open; open sites are initially empty.
A full site is an open site that can be connected to an open site in the top row via a chain of neighboring (left, right, up, down) open sites.
If there is a full site in the bottom row, then we say that the system percolates.
*/
