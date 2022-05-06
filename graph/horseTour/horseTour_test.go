package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenerateGraph(t *testing.T) {
	horseTourGraph := make(map[string][]string)

	nodeName := func(x, y int) string {
		return fmt.Sprintf("%c%d", (x + 'A'), (y + 1))
	}

	addNeighbor := func(x1, y1, x2, y2 int) {
		if 0 <= x2 && x2 < 8 && 0 <= y2 && y2 < 8 {
			nodeLabel := nodeName(x1, y1)
			nbrs := horseTourGraph[nodeLabel]
			nbrs = append(nbrs, nodeName(x2, y2))
			horseTourGraph[nodeLabel] = nbrs
		}
	}

	// iterate over nodes and collect neighbors
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			addNeighbor(x, y, x-1, y+2)
			addNeighbor(x, y, x-2, y+1)
			addNeighbor(x, y, x-2, y-1)
			addNeighbor(x, y, x-1, y-2)
			addNeighbor(x, y, x+1, y-2)
			addNeighbor(x, y, x+2, y-1)
			addNeighbor(x, y, x+2, y+1)
			addNeighbor(x, y, x+1, y+2)
		}
	}

	var sb strings.Builder
	sb.WriteString("strict graph {\n")

	for node, nbrs := range horseTourGraph {
		for _, nbr := range nbrs {
			sb.WriteString(node + "--" + nbr + "\n")
		}
	}

	sb.WriteString("}")

	fmt.Printf("\n%s", sb.String())
}
