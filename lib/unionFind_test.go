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
