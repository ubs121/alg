package common

import (
	"fmt"
	"testing"
)

// N ширхэг тоогоор зохиох бүх боломжит сэлгэмэл
func PermuteInts(nums []int) [][]int {
	n := len(nums)

	var ret [][]int

	// output initial state
	cmb := make([]int, n)
	copy(cmb, nums)
	ret = append(ret, cmb)

	// an encoding of the stack state.
	p := make([]int, n)

	i := 1
	for i < n {
		if p[i] < i {
			j := 0
			if i%2 > 0 {
				j = p[i]
			}

			// swap
			nums[j], nums[i] = nums[i], nums[j]

			// output
			cmb := make([]int, n)
			copy(cmb, nums)
			ret = append(ret, cmb)

			p[i]++
			i = 1
		} else {
			p[i] = 0
			i++
		}
	}
	return ret
}

func TestPermuteInts(t *testing.T) {
	a := []int{1, 2, 3} /* сэлгэх утгуудыг агуулах массив */
	ret := PermuteInts(a)
	fmt.Printf("%v", ret)
}

// recursive version
func permuteRecurse(k int, nums []int, out chan []int) {
	if k == 1 {
		// output
		cmb := make([]int, len(nums))
		copy(cmb, nums)
		out <- cmb
	} else {
		permuteRecurse(k-1, nums, out)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				nums[i], nums[k-1] = nums[k-1], nums[i]
			} else {
				nums[0], nums[k-1] = nums[k-1], nums[0]
			}

			permuteRecurse(k-1, nums, out)
		}
	}
}

func TestPermute2(t *testing.T) {
	a := []int{1, 2, 3} /* сэлгэх утгуудыг агуулах массив */

	out := make(chan []int)
	go func() {
		defer close(out)
		permuteRecurse(len(a), a, out)
	}()

	for v := range out {
		fmt.Printf("%v\n", v)
	}
}
