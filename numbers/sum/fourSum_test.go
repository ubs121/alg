package main

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	// 	1. Divide by half
	half := target / 2
	lHalf := map[int][][]int{}
	rHalf := map[int][][]int{}

	// calculate two sums
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := nums[i] + nums[j]
			if sum <= half {
				lHalf[sum] = append(lHalf[sum], []int{i, j})
			} else {
				rHalf[sum] = append(rHalf[sum], []int{i, j})
			}
		}
	}

	var ret [][]int
	dups := map[string]bool{}

	add := func(a, b, c, d int) {
		if b < c {
			lst := []int{nums[a], nums[b], nums[c], nums[d]}
			//key := (lst[0] ^ lst[1]) ^ (lst[2] ^ lst[3])
			key := fmt.Sprintf("%d/%d/%d/%d", lst[0], lst[1], lst[2], lst[3])
			if _, exists := dups[key]; !exists {
				ret = append(ret, lst)
				dups[key] = true
			}
		}
	}

	// Half pairs (only applies if target is even)
	if target%2 == 0 {
		if pairs, exists := lHalf[half]; exists {
			for i := 0; i < len(pairs)-1; i++ {
				a, b := pairs[i][0], pairs[i][1]
				for j := i + 1; j < len(pairs); j++ {
					c, d := pairs[j][0], pairs[j][1]
					add(a, b, c, d)
				}
			}
			delete(lHalf, half)
		}
	}

	// 2. Collect matches
	for k, pairs1 := range lHalf {
		if pairs2, exists := rHalf[target-k]; exists {
			for _, pr1 := range pairs1 {
				for _, pr2 := range pairs2 {
					add(pr1[0], pr1[1], pr2[0], pr2[1])
				}
			}
		}
	}

	return ret
}

func Test4Sum(t *testing.T) {
	// got := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	// fmt.Printf("%v\n", got)

	// got := fourSum([]int{2, 2, 2, 2, 2}, 8)
	// fmt.Printf("%v\n", got)

	// got := fourSum([]int{-3, -1, 0, 2, 4, 5}, 0)
	// fmt.Printf("%v\n", got)

	// got := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	// fmt.Printf("%v\n", got)

	// got := fourSum([]int{-1, 0, 1, 2, -1, -4}, -1) // [[-4,0,1,2],[-1,-1,0,1]]
	// fmt.Printf("%v\n", got)

	got := fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)
	fmt.Printf("%v\n", got)
	// [[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

	// got := fourSum([]int{-4, -3, -2, -1, 0, 0, 1, 2, 3, 4}, 0)
	// fmt.Printf("%v\n", got)
}

func Test4SumHash(t *testing.T) {
	hash := func(a, b, c, d int) int {
		return (a ^ b) ^ (c ^ d) + (a ^ c) ^ (b ^ d)
	}

	if hash(-1, 0, 0, 1) == hash(-3, 1, 0, 2) {
		t.Errorf("wrong hash")
	}
}

func Benchmark4Sum(b *testing.B) {
	//  81584	     14757 ns/op	    6539 B/op	     184 allocs/op
	// 173734	      7368 ns/op	    4632 B/op	     129 allocs/op
	for i := 0; i < b.N; i++ {
		fourSum([]int{-4, -3, -2, -1, 0, 0, 1, 2, 3, 4}, 0)
		// fmt.Printf("%v\n", got)
	}
}
