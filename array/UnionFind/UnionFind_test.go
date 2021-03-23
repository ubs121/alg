package graph

import "testing"

func TestUnion(t *testing.T) {

	uf := MakeUF(10)

	uf.Union(1, 2)
	uf.Union(2, 3)

}
