package alg

import (
	"fmt"
	"testing"
)

func TestGraphCreate(t *testing.T) {
	g := sampleGraph()
	exp := 3
	got := g.Degree(0)
	if got != exp {
		t.Errorf("exp %d, got %d", exp, got)
	}
}

func TestBFS(t *testing.T) {
	g := sampleGraph()

	// new bfs
	bfs := NewBFS(g, 2)
	if !bfs.HasPathTo(3) {
		t.Errorf("exp true, got false")
	}

	path := bfs.PathTo(3)
	fmt.Printf("2.pathTo(3) = %v", path)
}

func sampleGraph() *Graph {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(1, 3)
	return g
}
