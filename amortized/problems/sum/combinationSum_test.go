package main

import (
	"fmt"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	solver := newCombSumSolver(target, candidates, false)
	return solver.run()
}

func combinationSum2(candidates []int, target int) [][]int {
	solver := newCombSumSolver(target, candidates, true)
	return solver.run()
}

// combo sum solver
type comboSumSolver struct {
	target     int         // original target
	candidates []int       // original candidates
	limits     map[int]int // limits for each candidate, no entry means no limit
	out        chan []int  // output channel
	picks      map[int]int // current picks
}

func newCombSumSolver(target int, cs []int, withLimit bool) *comboSumSolver {
	sort.Ints(cs) // sort it

	sr := &comboSumSolver{}
	sr.target = target
	sr.candidates = cs
	sr.limits = make(map[int]int)
	sr.out = make(chan []int)
	sr.picks = make(map[int]int)

	if withLimit {
		// count candidates
		for _, c := range cs {
			sr.limits[c]++
		}
		// make it unique candidates
		sr.candidates = nil
		for c := range sr.limits {
			sr.candidates = append(sr.candidates, c)
		}
	}

	return sr
}

func (sr *comboSumSolver) run() [][]int {
	// solver
	go func() {
		defer close(sr.out)
		sr.solve(sr.target, sr.candidates)
	}()

	// collector
	var ret [][]int
	for {
		sol := <-sr.out
		if sol == nil {
			break
		}
		ret = append(ret, sol)
	}
	return ret
}

func (sr *comboSumSolver) solve(target int, candidates []int) {
	if target < 0 {
		return // invalid
	}

	if target == 0 { // one solution is found

		// translate it into a required format
		var sol []int
		for c, freq := range sr.picks {
			for i := 1; i <= freq; i++ {
				sol = append(sol, c)
			}
		}
		sr.out <- sol
		return
	}

	// check all candidates one by one
	for i, c := range candidates {
		limit := target / c

		if l, exists := sr.limits[c]; exists {
			limit = l // replace by the actual limit
		}

		// check multipliers of 'c'
		for k := 1; k <= limit; k++ {
			sr.picks[c] = k // pick 'c' k times
			sr.solve(target-c*k, candidates[i+1:])
		}
		sr.picks[c] = 0 // drop 'c'
	}
}

func TestComboSum(t *testing.T) {
	// ret := combinationSum([]int{2, 3, 5}, 8)
	// fmt.Printf("ret1=%v\n", ret)

	// ret2 := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	// fmt.Printf("ret2=%v", ret2)

	ret2 := combinationSum2([]int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}, 27)
	fmt.Printf("ret2=%v", ret2)
}
