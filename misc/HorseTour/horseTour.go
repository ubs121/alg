// Шатрын мориор нүд бүхэн дээр нэг удаа бууж хөлгийг бүтэн тойрох, Kind of Hamiltonian cycle
package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"

	alg "alg/lib"
	crypto_rand "crypto/rand"
)

// tour state
type tourState struct {
	tourGraph *alg.Graph  // graph of valid moves
	source    int         // starting node
	moves     map[int]int // moves completed (node->order)
}

// starts the horse tour from src
func start(src int) *tourState {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed with crypto rand")
	}
	crypto_rand.Read(b[:])
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	st := tourState{}
	st.source = src
	st.moves = make(map[int]int)
	st.buildTourGraph()
	st.tour(src)

	return &st
}

// check if tour completed
func (st *tourState) isGoal() bool {
	// TODO: check if it's a full cycle currentNode <-> firstNode
	return len(st.moves) == 64
}

// make tour starting from 'v'
func (st *tourState) tour(v int) bool {
	st.moves[v] = len(st.moves) // add into the completed moves

	if st.isGoal() {
		st.print() // print the result
		return true
	}

	nbrs := st.tourGraph.Adj[v] // get neighbors of 'v'
	mvs := make([]int, len(nbrs))
	copy(mvs, nbrs) // copy to preserve original moves

	solved := false
	for len(mvs) > 0 && !solved {
		ind := rand.Intn(len(mvs))              // select random move from 'nbrs'
		m := mvs[ind]                           // nbr id
		mvs = append(mvs[:ind], mvs[ind+1:]...) // remove this nbr

		// check if it's not a cycle
		if _, exists := st.moves[m]; !exists {
			solved = st.tour(m) // make the move
		}
	}

	// fixed moves
	// for _, m := range mvs {
	// 	if _, exists := st.moves[m]; !exists {
	// 		solved = st.nextMove(m)
	// 	}
	// 	if solved {
	// 		break
	// 	}
	// }

	delete(st.moves, v) // revert this node
	return solved
}

// build a tour graph, possible roads
func (st *tourState) buildTourGraph() {
	st.tourGraph = alg.NewGraph(64) // with 64 vertices

	// valid moves between vertices
	var validMoves = [][]int{{-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}}

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {

			v := x*8 + y
			for _, m := range validMoves {
				x1 := x + m[0]
				y1 := y + m[1]
				if 0 <= x1 && x1 < 8 && 0 <= y1 && y1 < 8 {
					st.tourGraph.AddEdge(v, (x1*8 + y1)) // add an edge
				}
			}
		}
	}
}

func (st *tourState) print() {
	var board [8][8]int
	for k, v := range st.moves {
		x := k / 8
		y := k % 8
		board[x][y] = v
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%2d ", board[i][j])
		}
		println()
	}
}

func main() {
	start(0) // start from (0,0)
}
