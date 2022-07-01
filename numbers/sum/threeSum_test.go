package main

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	zeros := 0
	neg := map[int]int{}
	pos := map[int]int{}
	for _, n := range nums {
		if n > 0 {
			pos[n]++
		} else if n < 0 {
			neg[n]++
		} else {
			zeros++
		}
	}

	ret := [][]int{}

	// zeros
	if zeros >= 3 {
		// add just one instance of (0,0,0)
		ret = append(ret, []int{0, 0, 0})
	}

	if len(neg) > 0 && len(pos) > 0 {
		negNums := getKeys(neg)
		posNums := getKeys(pos)

		x := 0 // 1st num
		y := 0 // 2nd num
		z := 0 // 4rd num z=x+y

		negMax := negNums[0]
		posMax := posNums[len(posNums)-1]

		// (--,+) pairs
		for i := len(negNums) - 1; i >= 0; i-- {
			x = negNums[i]
			if v := neg[x]; v > 1 {
				y = 2 * (-x)
				if _, exists := pos[y]; exists {
					ret = append(ret, []int{x, x, y}) // doubled number
				}
			}
			for j := i - 1; j >= 0; j-- {
				y = negNums[j]
				z = -(x + y)
				if posMax < z {
					break
				}
				if _, exists := pos[z]; exists {
					ret = append(ret, []int{y, x, z})
				}
			}
		}

		// (++,-) pairs
		for i := 0; i < len(posNums); i++ {
			x = posNums[i]
			if v := pos[x]; v > 1 {
				y = 2 * (-x)
				if _, exists := neg[y]; exists {
					ret = append(ret, []int{x, x, y}) // doubled number
				}
			}
			for j := i + 1; j < len(posNums); j++ {
				y = posNums[j]
				z = -(x + y)
				if z < negMax {
					break
				}
				if _, exists := neg[z]; exists {
					ret = append(ret, []int{z, x, y})
				}
			}
		}

		if zeros > 0 {
			// (-,0,+) pairs
			for i := 0; i < len(negNums); i++ {
				x = negNums[i]
				if _, exists := pos[-x]; exists {
					ret = append(ret, []int{x, 0, -x})
				}
			}
		}
	}
	return ret
}

func getKeys(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func Test3Sum(t *testing.T) {
	//ret := threeSum([]int{-1, 0, 1, 2, -1, -4})
	//ret := threeSum([]int{0, 0, 0, 0})
	// in: [-1,0,1,2,-1,-4,-2,-3,3,0,4]
	ret := threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})
	//exp: [[-4,0,4],[-3,0,3],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]   [-4,1,3],[-3,-1,4],[-3,1,2],
	fmt.Printf("%v", ret)
}
