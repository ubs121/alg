package main

import "sort"

func solveByMaxSplit(arr []int) int {
	if len(arr) < 2 {
		return 0 // no pair possible
	}

	maxPos := indexMax(arr) // max position, divide by this position
	max := arr[maxPos]      // max value
	left := arr[:maxPos]    // left array to the max
	right := arr[maxPos+1:] // right array to the max

	ch := make(chan int)
	defer close(ch)

	go func(a []int) {
		ch <- solveByMaxSplit(a)
	}(left)

	go func(a []int) {
		ch <- solveByMaxSplit(a)
	}(right)

	// count pairs crossing the max
	c := _countPairs(left, right, max)
	c += <-ch
	c += <-ch

	return c
}

// l - long array
// r - short array
func _countPairs(left, right []int, max int) int {
	l := make([]int, len(left))
	r := make([]int, len(right))
	copy(l, left)
	copy(r, right)

	sort.Ints(l)
	sort.Ints(r)

	c := 0          // number of valid pairs
	i := 0          // left index
	j := len(r) - 1 // right index

	for i < len(l) {
		x := l[i]

		// lower the right limit
		jj := j
		for jj >= 0 && max < x*r[jj] {
			jj-- // skip all large numbers
		}

		if jj < 0 {
			break // no number left on the right
		}

		c += len(r[:jj+1]) // +1 is to include r[j]
		j = jj
		i++
	}

	c += _count1(l) // count (1, max) pairs
	c += _count1(r) // count (max, 1) pairs

	return c
}

// a[] has to be sorted
func _count1(a []int) int {
	c := 0
	for i := 0; i < len(a) && a[i] == 1; i++ {
		c++
	}
	return c
}

// find index of max value, pick middle if there are many
func indexMax(a []int) int {
	maxVal := a[0]
	maxPos := 0
	var maxes []int
	for i := 1; i < len(a); i++ {
		if a[i] < maxVal {
			// nothing
		} else {
			if maxVal < a[i] {
				maxVal = a[i]
				maxPos = i
				maxes = nil // reset
			}
			// all 'i' where max<=a[i]
			maxes = append(maxes, i) // collect max position
		}
	}

	// take middle one if there are many 'max's
	if len(maxes) > 0 {
		maxPos = maxes[len(maxes)/2]
	}
	return maxPos
}
