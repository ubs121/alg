package main

import "alg/graph"

func main() {

	uf := graph.MakeUF(10)

	uf.Union(1, 2)
	uf.Union(2, 3)

	uf.Print()
}
